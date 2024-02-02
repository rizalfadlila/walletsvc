package deps

import (
	"github.com/julo/walletsvc/internal/domain/usecase"
)

type Service struct {
	WalletUC usecase.WalletUsecase
}
