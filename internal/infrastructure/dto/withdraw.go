package dto

import (
	"github.com/julo/walletsvc/internal/pkg/datatype"
)

type (
	WithdrawRequest struct {
		Amount      float64 `schema:"amount"`
		ReferenceID string  `schema:"reference_id"`
	}
)

type (
	WithdrawResponse struct {
		Id          string              `json:"id"`
		WithdrawnBy string              `json:"withdrawn_by"`
		Status      string              `json:"status"`
		WithdrawnAt datatype.SqlIsoTime `json:"withdrawn_at"`
		Amount      string              `json:"amount"`
		ReferenceId string              `json:"reference_id"`
	}
)
