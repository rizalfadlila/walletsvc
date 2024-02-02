package deps

import (
	"github.com/jmoiron/sqlx"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
	"github.com/julo/walletsvc/internal/infrastructure/router"
	"github.com/redis/go-redis/v9"
)

type Basic struct {
	Postgres *sqlx.DB
	Auth     auth.Auth
	Router   *router.MyRouter
	Redis    *redis.Client
}
