package wallet

import (
	"github.com/julo/walletsvc/internal/domain/repository"
	"github.com/julo/walletsvc/internal/domain/usecase"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
	"github.com/julo/walletsvc/internal/infrastructure/cache"
)

type (
	module struct {
		auth       auth.Auth
		cache      cache.Cache
		walletRepo repository.WalletRepository
	}

	Opts struct {
		Auth       auth.Auth
		Cache      cache.Cache
		WalletRepo repository.WalletRepository
	}
)

func New(o Opts) usecase.WalletUsecase {
	return &module{
		auth:       o.Auth,
		walletRepo: o.WalletRepo,
		cache:      o.Cache,
	}
}
