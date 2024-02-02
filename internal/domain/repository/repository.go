package repository

import (
	"context"
	"github.com/julo/walletsvc/internal/domain/entity"
	"github.com/julo/walletsvc/internal/domain/enums"
)

type (
	TransactionFunc func(ctx context.Context) error

	TransactionRepository interface {
		WithTransaction(ctx context.Context, fn TransactionFunc) error
	}

	WalletRepository interface {
		TransactionRepository
		Store(ctx context.Context, entity entity.Wallet) error
		UpdateStatusByUserID(ctx context.Context, status enums.StatusWallet, epocMilli int64, userID string) error
		UpdateBalance(ctx context.Context, entity entity.Wallet) error

		FindByUserID(ctx context.Context, userID string) (*entity.Wallet, error)
	}

	WalletTransactionRepository interface {
		Store(ctx context.Context, transaction entity.Transaction) error
	}
)
