package dto

import (
	"github.com/julo/walletsvc/internal/pkg/datatype"
)

type (
	DepositRequest struct {
		Amount      float64 `schema:"amount"`
		ReferenceID string  `schema:"reference_id"`
	}
)

type (
	DepositResponse struct {
		ID          string              `json:"id"`
		DepositedBy string              `json:"deposited_by"`
		Status      string              `json:"status"`
		DepositedAt datatype.SqlIsoTime `json:"deposited_at"`
		Amount      float64             `json:"amount"`
		ReferenceId string              `json:"reference_id"`
	}
)
