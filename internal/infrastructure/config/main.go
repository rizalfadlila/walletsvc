package config

import "github.com/julo/walletsvc/internal/infrastructure/auth"

type (
	MainConfig struct {
		Server Server      `envPrefix:"SERVER_"`
		DB     Database    `envPrefix:"DB_"`
		Auth   auth.Config `envPrefix:"AUTH_"`
		Redis  Redis       `envPrefix:"REDIS_"`
	}
)
