package store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
)

type CheckoutInput struct {
  UserID  int64
  Address map[string]any
  Courier string
  Service string
  Note    *string
}

type CheckoutResult struct {
  CheckoutGroup string   `json:"checkout_group_id"`
  OrderCodes    []string `json:"order_codes"`
}

func (db *DB) Checkout(ctx context.Context, in CheckoutInput) (*CheckoutResult, error) {
  tx, err := db.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable}); if err!=nil { return nil, err }
  defer tx.Rollback(ctx)

  // ambil item keranjang + group by shop
  rows, err := tx.Query(ctx, `
    SELECT p.shop_id, ci.product_id, p.title, ci.qty, ci.price_snapshot
    FROM carts c
    JOIN cart_items ci ON ci.cart_id=c.id
    JOIN products p ON p.id=ci.product_id
    WHERE c.user_id=$1
    ORDER BY p.shop_id`, in.UserID)
  if err!=nil { return nil, err }
  type item struct{ ShopID int64; ProductID int64; Title string; Qty int32; Price float64 }
  var items []item
  for rows.Next(){
    var it item
    if err:=rows.Scan(&it.ShopID,&it.ProductID,&it.Title,&it.Qty,&it.Price); err!=nil { return nil,err }
    items = append(items, it)
  }
  rows.Close()
  if len(items)==0 { return nil, fmt.Errorf("cart empty") }

  // lock & cek stok
  for _, it := range items {
    var stock int32
    if err := tx.QueryRow(ctx, `SELECT stock FROM products WHERE id=$1 FOR UPDATE`, it.ProductID).Scan(&stock); err!=nil { return nil, err }
    if stock < it.Qty { return nil, fmt.Errorf("stock not enough for product %d", it.ProductID) }
  }

  // buat orders per shop
  groupID := uuid.New().String()
  addrJSON, _ := json.Marshal(in.Address)

  // agregasi subtotal per shop
  type agg struct{ Subtotal float64; Items []item }
  perShop := map[int64]*agg{}
  for _, it := range items {
    a := perShop[it.ShopID]; if a==nil { a=&agg{}; perShop[it.ShopID]=a }
    a.Subtotal += float64(it.Qty)*it.Price
    a.Items = append(a.Items, it)
  }

  var codes []string
  for shopID, a := range perShop {
    code := fmt.Sprintf("ORD-%d-%d", shopID, time.Now().UnixNano())
    var orderID int64
    if err := tx.QueryRow(ctx, `
      INSERT INTO orders
      (code, checkout_group_id, user_id, shop_id, address_snapshot, status, payment_status, subtotal, shipping_cost, total, note)
      VALUES ($1,$2,$3,$4,$5,'pending','unpaid',$6,0,$6,$7)
      RETURNING id`,
      code, groupID, in.UserID, shopID, addrJSON, a.Subtotal, in.Note).Scan(&orderID); err!=nil { return nil, err }

    // order_items & decrement stok
    for _, it := range a.Items {
      _, err = tx.Exec(ctx, `
        INSERT INTO order_items(order_id, product_id, title_snapshot, price, qty)
        VALUES($1,$2,$3,$4,$5)`, orderID, it.ProductID, it.Title, it.Price, it.Qty)
      if err!=nil { return nil, err }
      _, err = tx.Exec(ctx, `UPDATE products SET stock=stock-$1 WHERE id=$2`, it.Qty, it.ProductID)
      if err!=nil { return nil, err }
    }
    codes = append(codes, code)
  }

  // kosongkan keranjang
  _, _ = tx.Exec(ctx, `
    DELETE FROM cart_items USING carts
    WHERE cart_items.cart_id=carts.id AND carts.user_id=$1`, in.UserID)

  if err := tx.Commit(ctx); err!=nil { return nil, err }
  return &CheckoutResult{CheckoutGroup: groupID, OrderCodes: codes}, nil
}
