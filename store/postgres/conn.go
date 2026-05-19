package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tom-kipp/elasticstore/store"
)

type (
	conn struct {
		ctx  context.Context
		conn *pgxpool.Conn
	}

	tx struct {
		ctx context.Context
		tx  pgx.Tx
	}
)

var (
	_ store.Conn = conn{}
	_ store.Tx   = tx{}
)
