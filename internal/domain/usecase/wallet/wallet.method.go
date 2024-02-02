package wallet

import (
	"context"
	"database/sql"
	"errors"
	"github.com/julo/walletsvc/internal/domain/constant"
	"github.com/julo/walletsvc/internal/domain/entity"
	"github.com/julo/walletsvc/internal/domain/enums"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/julo/walletsvc/internal/pkg/apierr"
	"github.com/julo/walletsvc/internal/pkg/ctxutil"
	"github.com/julo/walletsvc/internal/pkg/datatype"
	"net/http"
	"time"
)

func (u *module) RequestAccessAccount(ctx context.Context, request dto.RequestAccessAccount) (*dto.AccessAccount, *apierr.ApiErr) {
	wallet, err := u.walletRepo.FindByUserID(ctx, request.CustomerID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.New(constant.ErrQueryRequest.Error(), apierr.WithStacktrace(err))
	}

	token, err := u.auth.GenerateToken(auth.Claims{CustomerXID: request.CustomerID})
	if err != nil {
		return nil, apierr.New("error generating token", apierr.WithStacktrace(err))
	}

	accessToken := &dto.AccessAccount{Token: token}

	if wallet != nil {
		return accessToken, nil
	}

	if err = u.walletRepo.Store(ctx, entity.NewWallet(request.CustomerID)); err != nil {
		return nil, apierr.New(constant.ErrTransactionalRequest.Error(), apierr.WithStacktrace(err))
	}

	return accessToken, nil
}

func (u *module) EnableWallet(ctx context.Context, status enums.StatusWallet) (*dto.DetailWallet, *apierr.ApiErr) {
	var (
		claims = ctxutil.GetTokenClaim(ctx)
		wallet *dto.DetailWallet
	)

	err := u.walletRepo.WithTransaction(ctx, func(ctx context.Context) error {
		epocMilli := time.Now().Local().UnixMilli()

		result, err := u.walletRepo.FindByUserID(ctx, claims.CustomerXID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return apierr.New("error finding user",
					apierr.WithStatusCode(http.StatusNotFound),
					apierr.WithStacktrace(datatype.Fields{"error": "account not found"}))
			}

			return apierr.New(constant.ErrQueryRequest.Error(), apierr.WithStacktrace(err))
		}

		if result.EnabledAt != "" {
			return apierr.New(constant.ErrAlreadyEnabledWallet.Error(),
				apierr.WithStatusCode(http.StatusBadRequest),
				apierr.WithStacktrace(datatype.Fields{"error": constant.ErrAlreadyEnabledWallet.Error()}))
		}

		if err := u.walletRepo.UpdateStatusByUserID(ctx, status, epocMilli, claims.CustomerXID); err != nil {
			return apierr.New(constant.ErrTransactionalRequest.Error(), apierr.WithStacktrace(err))
		}

		wallet = result.ToDTODetailWallet()

		wallet.EnabledAt = datatype.SqlIsoTime(time.UnixMilli(epocMilli).Format(time.RFC3339))

		return nil
	})

	if err != nil {
		return nil, err.(*apierr.ApiErr)
	}

	return wallet, nil
}

func (u *module) ViewWallet(ctx context.Context) (*dto.DetailWallet, *apierr.ApiErr) {
	claims := ctxutil.GetTokenClaim(ctx)

	var (
		wallet dto.DetailWallet
	)

	if err := u.cache.GetWallet(ctx, claims.CustomerXID, &wallet); err == nil {
		return &wallet, nil
	}

	result, err := u.walletRepo.FindByUserID(ctx, claims.CustomerXID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierr.New("account not found", apierr.WithStacktrace(err), apierr.WithStatusCode(http.StatusNotFound))
		}
		return nil, apierr.New(constant.ErrQueryRequest.Error(), apierr.WithStacktrace(err))
	}

	if result.EnabledAt == "" {
		return nil, apierr.New(constant.ErrAlreadyEnabledWallet.Error(),
			apierr.WithStacktrace(datatype.Fields{"error": "Wallet disabled"}),
			apierr.WithStatusCode(http.StatusBadRequest))
	}

	b, err := result.Marshal()
	if err != nil {
		return nil, apierr.New("error marshaling data", apierr.WithStacktrace(err))
	}

	if err = u.cache.SetWallet(ctx, claims.CustomerXID, b); err != nil {
		return nil, apierr.New("error caching wallet", apierr.WithStacktrace(err))
	}

	return result.ToDTODetailWallet(), nil
}
