package wallet

import (
	"github.com/jmoiron/sqlx"
	"github.com/julo/walletsvc/internal/domain/repository"
)

type (
	module struct {
		*repository.BaseRepository
	}

	Opts struct {
		DB *sqlx.DB
	}
)

func New(o Opts) repository.WalletRepository {
	return &module{
		BaseRepository: repository.NewBaseRepository(o.DB),
	}
}
