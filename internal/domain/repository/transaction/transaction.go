package transaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/julo/walletsvc/internal/domain/repository"
)

type (
	module struct {
		*repository.BaseRepository
		db *sqlx.DB
	}

	Opts struct {
		DB *sqlx.DB
	}
)

func New(o Opts) repository.WalletTransactionRepository {
	return &module{
		BaseRepository: repository.NewBaseRepository(o.DB),
		db:             o.DB,
	}
}
