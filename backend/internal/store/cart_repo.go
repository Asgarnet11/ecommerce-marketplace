package store

import "context"

type CartRepo struct{ DB *DB }
func NewCartRepo(db *DB)*CartRepo{ return &CartRepo{DB:db} }

func (r *CartRepo) EnsureCart(ctx context.Context, userID int64) (int64, error) {
  var id int64
  err := r.DB.Pool.QueryRow(ctx, `
    INSERT INTO carts(user_id) VALUES($1)
    ON CONFLICT (user_id) DO UPDATE SET user_id=EXCLUDED.user_id
    RETURNING id`, userID).Scan(&id)
  return id, err
}

func (r *CartRepo) View(ctx context.Context, userID int64) ([]map[string]any, float64, error) {
  rows, err := r.DB.Pool.Query(ctx, `
    SELECT ci.product_id, p.title, ci.price_snapshot, ci.qty,
           (ci.price_snapshot*ci.qty)::numeric AS subtotal
    FROM cart_items ci
    JOIN carts c ON c.id=ci.cart_id AND c.user_id=$1
    JOIN products p ON p.id=ci.product_id`, userID)
  if err!=nil { return nil,0,err }
  defer rows.Close()
  var out []map[string]any
  var total float64
  for rows.Next(){
    var pid int64; var title string; var price float64; var qty int32; var sub float64
    if err:=rows.Scan(&pid,&title,&price,&qty,&sub); err!=nil { return nil,0,err }
    out = append(out, map[string]any{
      "product_id": pid, "title": title, "price": price, "qty": qty, "subtotal": sub,
    })
    total += sub
  }
  return out, total, nil
}

func (r *CartRepo) UpsertItem(ctx context.Context, userID, productID int64, qty int32) error {
  // ambil harga terbaru & id cart
  tx, err := r.DB.Pool.Begin(ctx); if err!=nil { return err }
  defer tx.Rollback(ctx)

  var cartID int64
  if err := tx.QueryRow(ctx, `
    INSERT INTO carts(user_id) VALUES($1)
    ON CONFLICT (user_id) DO UPDATE SET user_id=EXCLUDED.user_id
    RETURNING id`, userID).Scan(&cartID); err!=nil { return err }

  var price float64
  if err := tx.QueryRow(ctx, `SELECT price FROM products WHERE id=$1 AND status='active'`, productID).Scan(&price); err!=nil { return err }

  // qty <=0 berarti hapus
  if qty <= 0 {
    _, err = tx.Exec(ctx, `DELETE FROM cart_items WHERE cart_id=$1 AND product_id=$2`, cartID, productID)
  } else {
    _, err = tx.Exec(ctx, `
      INSERT INTO cart_items(cart_id, product_id, qty, price_snapshot)
      VALUES($1,$2,$3,$4)
      ON CONFLICT (cart_id,product_id)
      DO UPDATE SET qty=EXCLUDED.qty, price_snapshot=EXCLUDED.price_snapshot`,
      cartID, productID, qty, price)
  }
  if err!=nil { return err }
  return tx.Commit(ctx)
}

func (r *CartRepo) Remove(ctx context.Context, userID, productID int64) error {
  _, err := r.DB.Pool.Exec(ctx, `
    DELETE FROM cart_items
    USING carts
    WHERE carts.id=cart_items.cart_id AND carts.user_id=$1 AND cart_items.product_id=$2`, userID, productID)
  return err
}
