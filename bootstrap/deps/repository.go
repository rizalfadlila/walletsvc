package deps

import "github.com/julo/walletsvc/internal/domain/repository"

type Repository struct {
	Wallet repository.WalletRepository
}
