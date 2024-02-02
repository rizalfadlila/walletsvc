package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julo/walletsvc/internal/domain/enums"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/julo/walletsvc/internal/pkg/apierr"
	"github.com/julo/walletsvc/internal/pkg/datatype"
	"net/http"
)

func (a *API) CreateWalletAccount(ctx *fiber.Ctx) *dto.JSONResponse {
	var request dto.RequestAccessAccount

	if err := ctx.BodyParser(&request); err != nil {
		errParser := apierr.New("error parsing body",
			apierr.WithStacktrace(datatype.Fields{"request_access_account": err}),
			apierr.WithStatusCode(http.StatusBadRequest),
		)
		return dto.NewJSONResponse().WithError(errParser)

	}

	if err := request.Validate(); err != nil {
		errParser := apierr.New("error validating request",
			apierr.WithStacktrace(err),
			apierr.WithStatusCode(http.StatusBadRequest),
		)
		return dto.NewJSONResponse().WithError(errParser)
	}

	result, err := a.service.WalletUC.RequestAccessAccount(ctx.UserContext(), request)
	if err != nil {
		return dto.NewJSONResponse().WithError(err)
	}

	return dto.NewJSONResponse().
		WithStatusCode(http.StatusCreated).
		WithData(result)
}

func (a *API) EnableWallet(ctx *fiber.Ctx) *dto.JSONResponse {
	result, err := a.service.WalletUC.EnableWallet(ctx.UserContext(), enums.Enabled)
	if err != nil {
		return dto.NewJSONResponse().WithError(err)
	}

	return dto.NewJSONResponse().
		WithStatusCode(http.StatusCreated).
		WithData(result)
}

func (a *API) ViewWallet(ctx *fiber.Ctx) *dto.JSONResponse {
	result, err := a.service.WalletUC.ViewWallet(ctx.UserContext())
	if err != nil {
		return dto.NewJSONResponse().WithError(err)
	}

	return dto.NewJSONResponse().
		WithData(result)
}
