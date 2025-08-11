package store

import (
	"context"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone *string `json:"phone,omitempty"`
	Role  string `json:"role"`
	Status string `json:"status"`
	PwHash string `json:"-"`
}

type UserRepo struct{ DB *DB }
func NewUserRepo(db *DB) *UserRepo { return &UserRepo{DB: db} }

func (r *UserRepo) ByEmail(ctx context.Context, email string) (*User, error) {
	row := r.DB.Pool.QueryRow(ctx, `
SELECT id, name, email, phone, role, status, password_hash
FROM users WHERE email=$1 LIMIT 1`, email)
	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Role, &u.Status, &u.PwHash); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) Create(ctx context.Context, name, email, role, pwHash string) (*User, error) {
	row := r.DB.Pool.QueryRow(ctx, `
INSERT INTO users(name, email, role, password_hash, status)
VALUES($1,$2,$3,$4,'active')
RETURNING id, name, email, role, status`, name, email, role, pwHash)
	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.Status); err != nil {
		return nil, err
	}
	return &u, nil
}
