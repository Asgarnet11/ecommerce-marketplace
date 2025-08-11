package models

type CartItem struct {
	ProductID int64   `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Qty       int32   `json:"qty"`
	Subtotal  float64 `json:"subtotal"`
}
