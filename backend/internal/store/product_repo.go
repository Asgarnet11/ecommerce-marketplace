package store

import (
	"context"
	"fmt"
	"strings"

	"marketplace/backend/internal/models"
)

type ProductRepo struct{ DB *DB }

func NewProductRepo(db *DB) *ProductRepo { return &ProductRepo{DB: db} }

type ListParams struct {
	Search string
	Sort   string // relevance|newest|price_asc|price_desc
	Page   int
	Per    int
}

func (r *ProductRepo) List(ctx context.Context, p ListParams) ([]models.Product, int, error) {
	if p.Page <= 0 { p.Page = 1 }
	if p.Per <= 0 || p.Per > 100 { p.Per = 20 }

	q := strings.TrimSpace(p.Search)

	base := `
SELECT p.id, p.shop_id, p.title, p.slug, COALESCE(p.description,''), p.price, p.stock,
       COALESCE((SELECT i.url FROM product_images i WHERE i.product_id=p.id AND i.is_primary=true LIMIT 1),'') AS image
FROM products p
WHERE p.status='active'`
	args := []any{}
	if q != "" {
		base += " AND (p.title ILIKE $1 OR p.slug ILIKE $1)"
		args = append(args, "%"+q+"%")
	}

	order := " ORDER BY p.id DESC"
	switch p.Sort {
	case "newest":
		order = " ORDER BY p.created_at DESC"
	case "price_asc":
		order = " ORDER BY p.price ASC"
	case "price_desc":
		order = " ORDER BY p.price DESC"
	}

	countSQL := "SELECT COUNT(*) FROM (" + base + ") x"
	var total int
	if err := r.DB.Pool.QueryRow(ctx, countSQL, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limitOffset := fmt.Sprintf(" LIMIT %d OFFSET %d", p.Per, (p.Page-1)*p.Per)
	sql := base + order + limitOffset

	rows, err := r.DB.Pool.Query(ctx, sql, args...)
	if err != nil { return nil, 0, err }
	defer rows.Close()

	items := []models.Product{}
	for rows.Next() {
		var it models.Product
		if err := rows.Scan(&it.ID, &it.ShopID, &it.Title, &it.Slug, &it.Desc, &it.Price, &it.Stock, &it.Image); err != nil {
			return nil, 0, err
		}
		items = append(items, it)
	}
	return items, total, nil
}

func (r *ProductRepo) BySlug(ctx context.Context, slug string) (*models.Product, error) {
	sql := `
SELECT p.id, p.shop_id, p.title, p.slug, COALESCE(p.description,''), p.price, p.stock,
       COALESCE((SELECT i.url FROM product_images i WHERE i.product_id=p.id AND i.is_primary=true LIMIT 1),'') AS image
FROM products p
WHERE p.slug=$1 AND p.status='active'
LIMIT 1`
	row := r.DB.Pool.QueryRow(ctx, sql, slug)
	var it models.Product
	if err := row.Scan(&it.ID, &it.ShopID, &it.Title, &it.Slug, &it.Desc, &it.Price, &it.Stock, &it.Image); err != nil {
		return nil, err
	}
	return &it, nil
}
