package config

import "github.com/julo/walletsvc/internal/infrastructure/router"

type (
	Server struct {
		GracefulTimeout string        `env:"GRACEFUL_TIMEOUT"`
		Rest            router.Config `envPrefix:"REST_"`
	}
)
