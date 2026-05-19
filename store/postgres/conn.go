package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

	pgConn interface {
		QueryRow(ctx context.Context, stmt string, args ...any) pgx.Row
		Query(ctx context.Context, stmt string, args ...any) (pgx.Rows, error)
		Exec(ctx context.Context, stmt string, args ...any) (pgconn.CommandTag, error)
	}
)

func (c conn) Context() context.Context {
	return c.ctx
}

func (c conn) Exists(pred store.Pred) (bool, error) {
	return exists(c.ctx, c.conn, pred)
}

func (c conn) Count(pred store.Pred) (uint64, error) {
	return count(c.ctx, c.conn, pred)
}

func (c conn) First(pred store.Pred) (*store.Object, error) {
	return first(c.ctx, c.conn, pred)
}

func (c conn) Distinct(path store.Path, pred store.Pred) ([]any, error) {
	return distinct(c.ctx, c.conn, path, pred)
}

func (c conn) All(pred store.Pred, order store.Ordering, offset uint64, limit uint64) ([]store.Object, error) {
	return all(c.ctx, c.conn, pred, order, offset, limit)
}

func (t tx) Context() context.Context {
	return t.ctx
}

func (t tx) Exists(pred store.Pred) (bool, error) {
	return exists(t.ctx, t.tx, pred)
}

func (t tx) Count(pred store.Pred) (uint64, error) {
	return count(t.ctx, t.tx, pred)
}

func (t tx) First(pred store.Pred) (*store.Object, error) {
	return first(t.ctx, t.tx, pred)
}

func (t tx) Distinct(path store.Path, pred store.Pred) ([]any, error) {
	return distinct(t.ctx, t.tx, path, pred)
}

func (t tx) All(pred store.Pred, order store.Ordering, offset uint64, limit uint64) ([]store.Object, error) {
	return all(t.ctx, t.tx, pred, order, offset, limit)
}

func exists(ctx context.Context, c pgConn, pred store.Pred) (bool, error) {

}

func count(ctx context.Context, c pgConn, pred store.Pred) (uint64, error) {

}

func first(ctx context.Context, c pgConn, pred store.Pred) (*store.Object, error) {

}

func distinct(ctx context.Context, c pgConn, path store.Path, pred store.Pred) ([]any, error) {

}

func all(ctx context.Context, c pgConn, pred store.Pred, order store.Ordering, offset uint64, limit uint64) ([]store.Object, error) {

}

var (
	_ store.Conn = conn{}
	_ store.Tx   = tx{}
	_ pgConn     = (*pgxpool.Conn)(nil)
	_ pgConn     = (pgx.Tx)(nil)
)
