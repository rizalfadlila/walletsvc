package dto

import (
	"github.com/julo/walletsvc/internal/pkg/datatype"
)

type (
	TransactionResponse struct {
		Amount       float64             `json:"amount"`
		TransactedAt datatype.SqlIsoTime `json:"transacted_at"`
		ReferenceID  string              `json:"reference_id"`
		ID           string              `json:"id"`
		Type         string              `json:"type"`
		Status       string              `json:"status"`
	}
)
