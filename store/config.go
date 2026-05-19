package store

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func Config(ctx context.Context) error {
	key, connStr, ok := strings.Cut(os.Getenv("STORE_CONN"), ":")
	if !ok {
		return fmt.Errorf("store connection variable 'STORE_CONN' does not have the '<type>:<conn_string>' format")
	}

	prov, ok := providers[key]
	if !ok {
		return fmt.Errorf("store provider '%s' is not registered", key)
	}

	p, err := prov(ctx, connStr)
	if err != nil {
		return fmt.Errorf("error creating store pool: %s", err.Error())
	}

	pool = p
	return nil
}

func Register(name string, fn Provider) {
	providers[name] = fn
}

type (
	Pool interface {
		Conn(ctx context.Context, f func(c Conn) error) error
		Tx(ctx context.Context, f func(t Tx) error) error
	}

	Provider func(ctx context.Context, connStr string) (Pool, error)
)

var (
	pool      Pool
	providers map[string]Provider = make(map[string]Provider)
)
