package initiator

import (
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/internal/domain/repository"
	"github.com/julo/walletsvc/internal/domain/repository/wallet"
)

func (i *Initiator) InitRepositories() *Initiator {
	i.repositories = &deps.Repository{
		Wallet: i.NewWalletRepository(),
	}

	return i
}

func (i *Initiator) NewWalletRepository() repository.WalletRepository {
	return wallet.New(wallet.Opts{
		DB: i.basic.Postgres,
	})
}
