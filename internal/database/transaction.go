package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Connection interface {
	DBTX
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

func Transaction[T any](ctx context.Context, conn Connection, options pgx.TxOptions, do func(context.Context, DBTX) (T, error)) (T, error) {
	var zero T
	tx, err := conn.BeginTx(ctx, options)
	if err != nil {
		return zero, err
	}
	result, err := do(ctx, tx)
	if err != nil {
		return zero, tx.Rollback(ctx)
	}
	if err := tx.Commit(ctx); err != nil {
		return zero, err
	}
	return result, nil
}
