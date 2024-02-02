package ctxutil

import "context"

func SetSqlTx(ctx context.Context, value interface{}) context.Context {
	return context.WithValue(ctx, keySQLTransaction, value)
}

func SetTokenClaim(ctx context.Context, value interface{}) context.Context {
	return context.WithValue(ctx, keyTokenClaim, value)
}
