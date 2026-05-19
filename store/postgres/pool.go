package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tom-kipp/elasticstore/store"
)

type (
	Pool struct {
		p *pgxpool.Pool
	}
)

func (p *Pool) Conn(ctx context.Context, f func(c store.Conn) error) error {
	return p.p.AcquireFunc(ctx, func(c *pgxpool.Conn) error {
		return f(conn{ctx: ctx, conn: c})
	})
}

func (p *Pool) Tx(ctx context.Context, f func(c store.Tx) error) error {
	ptx, err := p.p.Begin(ctx)
	if err != nil {
		return err
	}

	defer ptx.Rollback(ctx)

	if err := f(tx{ctx: ctx, tx: ptx}); err != nil {
		return err
	}

	return ptx.Commit(ctx)
}

var (
	_ store.Pool = (*Pool)(nil)
)
