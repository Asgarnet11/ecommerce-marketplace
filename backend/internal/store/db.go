package store

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool() *pgxpool.Pool {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" { log.Fatal("DB_DSN not set") }
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil { log.Fatal(err) }
	return pool
}
