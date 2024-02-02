package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/julo/walletsvc/internal/infrastructure/auth"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/julo/walletsvc/internal/pkg/apierr"
	"github.com/julo/walletsvc/internal/pkg/ctxutil"
	"github.com/julo/walletsvc/internal/pkg/datatype"
	"net/http"
	"strings"
)

var (
	errEmptyToken = apierr.New("unauthorized",
		apierr.WithStatusCode(http.StatusBadRequest),
		apierr.WithStacktrace(datatype.Fields{"token": "missing token"}),
	)
	errInvalidFormatToken = apierr.New("unauthorized",
		apierr.WithStatusCode(http.StatusBadRequest),
		apierr.WithStacktrace(datatype.Fields{"token": "invalid format"}),
	)
)

func NewJwt(auth auth.Auth) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		if authorization == "" {
			return dto.NewJSONResponse().WithError(errEmptyToken).Send(ctx)
		}

		token := strings.Split(authorization, " ")

		if len(token) != 2 {
			return dto.NewJSONResponse().WithError(errInvalidFormatToken).Send(ctx)
		}

		claim, err := auth.VerifyToken(token[1])
		if err != nil {
			errInvalidToken := apierr.New("unauthorized",
				apierr.WithStatusCode(http.StatusUnauthorized),
				apierr.WithStacktrace(
					datatype.Fields{"token": fmt.Sprintf("%v", err)}),
			)
			return dto.NewJSONResponse().WithError(errInvalidToken).Send(ctx)
		}

		ctx.SetUserContext(ctxutil.SetTokenClaim(ctx.UserContext(), claim))

		return ctx.Next()
	}
}
