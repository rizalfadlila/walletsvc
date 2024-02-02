package usecase

import (
	"context"
	"github.com/julo/walletsvc/internal/domain/enums"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/julo/walletsvc/internal/pkg/apierr"
)

type (
	WalletUsecase interface {
		RequestAccessAccount(ctx context.Context, request dto.RequestAccessAccount) (*dto.AccessAccount, *apierr.ApiErr)
		EnableWallet(ctx context.Context, status enums.StatusWallet) (*dto.DetailWallet, *apierr.ApiErr)
		ViewWallet(ctx context.Context) (*dto.DetailWallet, *apierr.ApiErr)
	}
)
