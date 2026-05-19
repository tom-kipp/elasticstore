package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tom-kipp/elasticstore/store"
)

func Provider(ctx context.Context, connStr string) (store.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	p, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Pool{p: p}, nil
}

func init() {
	store.Register("postgres", Provider)
}
