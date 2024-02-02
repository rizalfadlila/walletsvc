package initiator

import (
	"github.com/caarlos0/env/v6"
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/internal/infrastructure/config"
	"github.com/rs/zerolog/log"
)

type (
	Initiator struct {
		config       *config.MainConfig
		basic        *deps.Basic
		repositories *deps.Repository
		services     *deps.Service
	}
)

func New() *Initiator {
	return &Initiator{
		config: parseConfig(),
	}
}

func parseConfig() *config.MainConfig {
	cfg := &config.MainConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	return cfg
}

func (i *Initiator) SetupDeps() *deps.Deps {
	return &deps.Deps{
		Config:  i.config,
		Basic:   i.basic,
		Service: i.services,
	}
}
