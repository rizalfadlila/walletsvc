package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/julo/walletsvc/internal/pkg/ctxutil"
	"github.com/rs/zerolog/log"
)

type (
	BaseRepository struct {
		db *sqlx.DB
	}

	Execer interface {
		sqlx.Execer
		sqlx.ExecerContext
		NamedExec(query string, arg interface{}) (sql.Result, error)
		NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	}

	Queryer interface {
		sqlx.Queryer
		GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
		SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	}

	QueryExecer interface {
		Queryer
		Execer
	}
)

func NewBaseRepository(db *sqlx.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

func (b *BaseRepository) DB(ctx context.Context) QueryExecer {
	if tx := ctxutil.GetSqlTx(ctx); tx != nil {
		return tx
	}

	return b.db
}

func (b *BaseRepository) WithTransaction(ctx context.Context, fn TransactionFunc) error {
	if parentTx := ctxutil.GetSqlTx(ctx); parentTx != nil {
		return fn(ctx)
	}

	tx, err := b.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func(tx *sqlx.Tx) {
		if err = tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Err(err).Msg("failed on rollback transaction")
		}
	}(tx)

	if err = fn(ctxutil.SetSqlTx(ctx, tx)); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
