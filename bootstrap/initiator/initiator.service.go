package initiator

import (
	"github.com/julo/walletsvc/bootstrap/deps"
	"github.com/julo/walletsvc/internal/domain/usecase"
	"github.com/julo/walletsvc/internal/domain/usecase/wallet"
	"github.com/julo/walletsvc/internal/infrastructure/cache"
)

func (i *Initiator) InitService() *Initiator {
	i.services = &deps.Service{
		WalletUC: i.NewWalletUsecase(),
	}

	return i
}

func (i *Initiator) NewWalletUsecase() usecase.WalletUsecase {
	return wallet.New(wallet.Opts{
		Auth:       i.basic.Auth,
		WalletRepo: i.repositories.Wallet,
		Cache:      cache.New(i.basic.Redis),
	})
}
