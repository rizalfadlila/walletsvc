package initiator

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
	"github.com/julo/walletsvc/internal/infrastructure/router"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func (i *Initiator) InitBasic() *Initiator {
	i.basic = &deps.Basic{
		Router:   i.NewRouter(),
		Auth:     i.NewAuth(),
		Postgres: i.NewPostgresClient(),
		Redis:    i.NewCache(),
	}

	return i
}

func (i *Initiator) NewRouter() *router.MyRouter {
	r, err := router.New(i.config.Server.Rest)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to init router")
	}

	return r
}

func (i *Initiator) NewAuth() auth.Auth {
	return auth.New(i.config.Auth)
}

func (i *Initiator) NewPostgresClient() *sqlx.DB {
	cfg := i.config.DB

	masterDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", masterDSN)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open DB connection")
	}

	if err = db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping DB")
	}

	return db
}

func (i *Initiator) NewCache() *redis.Client {
	opts, err := redis.ParseURL(i.config.Redis.URL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init redis client")
	}

	client := redis.NewClient(opts)

	if err = client.Ping(context.TODO()).Err(); err != nil {
		log.Fatal().Err(err).Msg("failed to init redis client")
	}

	return client
}
