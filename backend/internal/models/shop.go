package models

type Shop struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Slug      string  `json:"slug"`
	LogoURL   *string `json:"logo_url,omitempty"`
	Status    string  `json:"status"`
	ProdCount int64   `json:"product_count,omitempty"`
}
