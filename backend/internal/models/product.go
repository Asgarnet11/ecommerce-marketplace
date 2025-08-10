package models

type Product struct {
	ID     int64   `json:"id"`
	ShopID int64   `json:"shop_id"`
	Title  string  `json:"title"`
	Slug   string  `json:"slug"`
	Desc   string  `json:"description"`
	Price  float64 `json:"price"`
	Stock  int32   `json:"stock"`
	Image  string  `json:"image"`
}
