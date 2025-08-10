package store

import (
	"context"
	"fmt"
	"marketplace/backend/internal/models"
)

type ShopRepo struct{ DB *DB }
func NewShopRepo(db *DB) *ShopRepo { return &ShopRepo{DB: db} }

type ShopListParams struct {
	Page int
	Per  int
}

func (r *ShopRepo) List(ctx context.Context, p ShopListParams) ([]models.Shop, int, error) {
	if p.Page <= 0 { p.Page = 1 }
	if p.Per <= 0 || p.Per > 100 { p.Per = 20 }

	var total int
	if err := r.DB.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM shops WHERE status='active'`).Scan(&total); err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf(`
SELECT s.id, s.name, s.slug, s.logo_url, s.status,
       (SELECT COUNT(*) FROM products pr WHERE pr.shop_id=s.id AND pr.status='active') AS product_count
FROM shops s
WHERE s.status='active'
ORDER BY s.name
LIMIT %d OFFSET %d`, p.Per, (p.Page-1)*p.Per)

	rows, err := r.DB.Pool.Query(ctx, sql)
	if err != nil { return nil, 0, err }
	defer rows.Close()

	var out []models.Shop
	for rows.Next() {
		var sh models.Shop
		if err := rows.Scan(&sh.ID, &sh.Name, &sh.Slug, &sh.LogoURL, &sh.Status, &sh.ProdCount); err != nil {
			return nil, 0, err
		}
		out = append(out, sh)
	}
	return out, total, nil
}

func (r *ShopRepo) BySlug(ctx context.Context, slug string) (*models.Shop, error) {
	row := r.DB.Pool.QueryRow(ctx, `
SELECT s.id, s.name, s.slug, s.logo_url, s.status,
       (SELECT COUNT(*) FROM products pr WHERE pr.shop_id=s.id AND pr.status='active') AS product_count
FROM shops s
WHERE s.slug=$1 AND s.status IN ('active','pending')
LIMIT 1`, slug)
	var sh models.Shop
	if err := row.Scan(&sh.ID, &sh.Name, &sh.Slug, &sh.LogoURL, &sh.Status, &sh.ProdCount); err != nil {
		return nil, err
	}
	return &sh, nil
}
