package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/julo/walletsvc/internal/pkg/apierr"
	"net/http"
	"runtime/debug"
)

func (r *MyRouter) convertHandler(handler Handler) fiber.Handler {
	h := func(ctx *fiber.Ctx) error {
		resp := handler(ctx)
		return resp.Send(ctx)
	}

	return h
}

func recovery() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				internalErr := apierr.New("internal server error",
					apierr.WithStacktrace(string(debug.Stack())),
					apierr.WithStatusCode(http.StatusInternalServerError),
				)

				err = dto.NewJSONResponse().
					WithError(internalErr).
					Send(ctx)
				return
			}
		}()

		return ctx.Next()
	}
}
