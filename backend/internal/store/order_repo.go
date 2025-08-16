package store

import (
	"context"
	"fmt"
)

type Order struct {
	ID        int64    `json:"id"`
	Code      string   `json:"code"`
	Status    string   `json:"status"`
	PayStatus string   `json:"payment_status"`
	Subtotal  float64  `json:"subtotal"`
	ShipCost  float64  `json:"shipping_cost"`
	Total     float64  `json:"total"`
	ShopID    int64    `json:"shop_id"`
	ShopName  string   `json:"shop_name"`
	CreatedAt string   `json:"created_at"`
}

type OrderItem struct {
	ProductID int64   `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Qty       int32   `json:"qty"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderDetail struct {
	Order
	AddressJSON map[string]any `json:"address"`
	Payment     struct {
		Provider string  `json:"provider"`
		Method   *string `json:"method"`
		Amount   float64 `json:"amount"`
		Status   string  `json:"status"`
	} `json:"payment"`
	Shipment struct {
		Courier  string  `json:"courier"`
		Service  *string `json:"service"`
		Tracking *string `json:"tracking_code"`
		Status   string  `json:"status"`
	} `json:"shipment"`
	Items []OrderItem `json:"items"`
}

type OrderRepo struct{ DB *DB }
func NewOrderRepo(db *DB) *OrderRepo { return &OrderRepo{DB: db} }

type OrderListParams struct {
	Status string
	Page   int
	Per    int
}

func (r *OrderRepo) ListByUser(ctx context.Context, userID int64, p OrderListParams) ([]Order, int, error) {
	if p.Page <= 0 { p.Page = 1 }
	if p.Per <= 0 || p.Per > 100 { p.Per = 10 }

	where := "o.user_id=$1"
	args := []any{userID}
	if p.Status != "" {
		where += " AND o.status=$2"
		args = append(args, p.Status)
	}

	countSQL := "SELECT COUNT(*) FROM orders o WHERE " + where
	var total int
	if err := r.DB.Pool.QueryRow(ctx, countSQL, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf(`
SELECT o.id, o.code, o.status, o.payment_status, o.subtotal, o.shipping_cost, o.total,
       o.shop_id, s.name AS shop_name, to_char(o.created_at,'YYYY-MM-DD HH24:MI:SS') AS created_at
FROM orders o
JOIN shops s ON s.id=o.shop_id
WHERE %s
ORDER BY o.created_at DESC
LIMIT %d OFFSET %d`, where, p.Per, (p.Page-1)*p.Per)

	rows, err := r.DB.Pool.Query(ctx, sql, args...)
	if err != nil { return nil, 0, err }
	defer rows.Close()

	var out []Order
	for rows.Next() {
		var it Order
		if err := rows.Scan(&it.ID, &it.Code, &it.Status, &it.PayStatus, &it.Subtotal, &it.ShipCost, &it.Total, &it.ShopID, &it.ShopName, &it.CreatedAt); err != nil {
			return nil, 0, err
		}
		out = append(out, it)
	}
	return out, total, nil
}

func (r *OrderRepo) DetailByCode(ctx context.Context, userID int64, code string) (*OrderDetail, error) {
	row := r.DB.Pool.QueryRow(ctx, `
SELECT o.id, o.code, o.status, o.payment_status, o.subtotal, o.shipping_cost, o.total,
       o.shop_id, s.name AS shop_name, to_char(o.created_at,'YYYY-MM-DD HH24:MI:SS') AS created_at,
       o.address_snapshot
FROM orders o
JOIN shops s ON s.id=o.shop_id
WHERE o.user_id=$1 AND o.code=$2
LIMIT 1`, userID, code)

	var d OrderDetail
	if err := row.Scan(&d.ID, &d.Code, &d.Status, &d.PayStatus, &d.Subtotal, &d.ShipCost, &d.Total,
		&d.ShopID, &d.ShopName, &d.CreatedAt, &d.AddressJSON); err != nil {
		return nil, err
	}
	// payment (optional)
	_ = r.DB.Pool.QueryRow(ctx, `
SELECT COALESCE(provider,''), method, COALESCE(amount,0), COALESCE(status,'pending')
FROM payments WHERE order_id=$1`, d.ID).
		Scan(&d.Payment.Provider, &d.Payment.Method, &d.Payment.Amount, &d.Payment.Status)

	// shipment (optional)
	_ = r.DB.Pool.QueryRow(ctx, `
SELECT COALESCE(courier,''), service, tracking_code, COALESCE(status,'pending')
FROM shipments WHERE order_id=$1`, d.ID).
		Scan(&d.Shipment.Courier, &d.Shipment.Service, &d.Shipment.Tracking, &d.Shipment.Status)

	// items
	rows, err := r.DB.Pool.Query(ctx, `
SELECT product_id, title_snapshot, price, qty, (price*qty)::numeric
FROM order_items WHERE order_id=$1 ORDER BY id`, d.ID)
	if err != nil { return nil, err }
	defer rows.Close()
	for rows.Next() {
		var it OrderItem
		if err := rows.Scan(&it.ProductID, &it.Title, &it.Price, &it.Qty, &it.Subtotal); err != nil {
			return nil, err
		}
		d.Items = append(d.Items, it)
	}
	return &d, nil
}
