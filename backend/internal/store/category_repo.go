package store

import (
	"context"
	"marketplace/backend/internal/models"
)

type CategoryRepo struct{ DB *DB }

func NewCategoryRepo(db *DB) *CategoryRepo { return &CategoryRepo{DB: db} }

// List returns flat categories
func (r *CategoryRepo) List(ctx context.Context) ([]models.Category, error) {
	rows, err := r.DB.Pool.Query(ctx, `
		SELECT id, parent_id, name, slug
		FROM categories
		ORDER BY COALESCE(parent_id, 0), sort_order, name`)
	if err != nil { return nil, err }
	defer rows.Close()

	var out []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.ParentID, &c.Name, &c.Slug); err != nil { return nil, err }
		out = append(out, c)
	}
	return out, nil
}

// ListTree builds a tree in-memory
func (r *CategoryRepo) ListTree(ctx context.Context) ([]models.Category, error) {
	all, err := r.List(ctx)
	if err != nil { return nil, err }
	byID := map[int64]*models.Category{}
	var roots []models.Category

	for i := range all {
		byID[all[i].ID] = &all[i]
	}
	for i := range all {
		c := &all[i]
		if c.ParentID == nil {
			roots = append(roots, *c)
		} else if p := byID[*c.ParentID]; p != nil {
			p.Children = append(p.Children, *c)
		}
	}
	return roots, nil
}
