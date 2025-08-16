package store

import (
	"context"
	"strings"
)

type SearchRepo struct{ DB *DB }
func NewSearchRepo(db *DB)*SearchRepo{ return &SearchRepo{DB: db} }

type SearchResultProduct struct {
	ID    int64   `json:"id"`
	ShopID int64  `json:"shop_id"`
	Title string  `json:"title"`
	Slug  string  `json:"slug"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

type SearchResultShop struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
}

type SearchResults struct {
	Products []SearchResultProduct `json:"products"`
	Shops    []SearchResultShop    `json:"shops"`
}

func (r *SearchRepo) Search(ctx context.Context, q string, limitProd, limitShop int) (*SearchResults, error) {
	q = strings.TrimSpace(q)
	if limitProd <= 0 || limitProd > 50 { limitProd = 10 }
	if limitShop <= 0 || limitShop > 50 { limitShop = 5 }

	// products
	prows, err := r.DB.Pool.Query(ctx, `
SELECT p.id, p.shop_id, p.title, p.slug,
       COALESCE((SELECT i.url FROM product_images i WHERE i.product_id=p.id AND i.is_primary=true LIMIT 1),'') AS image,
       p.price
FROM products p
WHERE p.status='active' AND (p.title ILIKE $1 OR p.slug ILIKE $1)
ORDER BY p.id DESC
LIMIT $2`, "%"+q+"%", limitProd)
	if err != nil { return nil, err }
	defer prows.Close()

	var prods []SearchResultProduct
	for prows.Next() {
		var it SearchResultProduct
		if err := prows.Scan(&it.ID,&it.ShopID,&it.Title,&it.Slug,&it.Image,&it.Price); err != nil { return nil, err }
		prods = append(prods, it)
	}

	// shops
	srows, err := r.DB.Pool.Query(ctx, `
SELECT s.id, s.name, s.slug
FROM shops s
WHERE s.status='active' AND (s.name ILIKE $1 OR s.slug ILIKE $1)
ORDER BY s.name
LIMIT $2`, "%"+q+"%", limitShop)
	if err != nil { return nil, err }
	defer srows.Close()

	var shops []SearchResultShop
	for srows.Next() {
		var it SearchResultShop
		if err := srows.Scan(&it.ID,&it.Name,&it.Slug); err != nil { return nil, err }
		shops = append(shops, it)
	}

	return &SearchResults{Products: prods, Shops: shops}, nil
}
