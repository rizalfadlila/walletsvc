package ctxutil

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
)

func GetSqlTx(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(keySQLTransaction).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}

func GetTokenClaim(ctx context.Context) *auth.Claims {
	if tx, ok := ctx.Value(keyTokenClaim).(*auth.Claims); ok {
		return tx
	}
	return nil
}
