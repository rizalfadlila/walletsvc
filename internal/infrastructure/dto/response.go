package dto

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/julo/walletsvc/internal/pkg/apierr"
	"github.com/rs/zerolog/log"
	"net/http"
)

type (
	// JSONResponse with structure based on JSend (https://github.com/omniti-labs/jsend).
	JSONResponse struct {
		HttpStatusCode int         `json:"-"`
		Status         string      `json:"status"`
		Data           interface{} `json:"data,omitempty"`
		Message        string      `json:"message,omitempty"`
	}
)

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{
		Status:         "success",
		HttpStatusCode: http.StatusOK,
	}
}

func (r *JSONResponse) WithData(d interface{}) *JSONResponse {
	r.Data = d
	return r
}

func (r *JSONResponse) WithStatusCode(code int) *JSONResponse {
	r.HttpStatusCode = code
	return r
}

func (r *JSONResponse) WithError(err *apierr.ApiErr) *JSONResponse {
	r.HttpStatusCode = err.StatusCode

	if r.HttpStatusCode >= 400 && r.HttpStatusCode < 500 {
		r.Data = err.Stacktrace
		r.Status = "fail"
	} else if r.HttpStatusCode >= 500 {
		r.Message = err.Message
		r.Status = "error"
	}

	log.Err(err).Str("stacktrace", fmt.Sprintf("%v", err.Stacktrace)).Msg("error api")

	return r
}

func (r *JSONResponse) Send(ctx *fiber.Ctx) error {
	return ctx.
		Status(r.HttpStatusCode).
		JSON(*r)
}
