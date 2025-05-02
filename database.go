package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func newDB(user, password, port, database string) (*DB, error) {
	conn := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		user, password, port, database,
	)
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	db := &DB{pool}
	return db, nil
}
