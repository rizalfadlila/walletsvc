package dto

import (
	"github.com/julo/walletsvc/internal/domain/enums"
	"github.com/julo/walletsvc/internal/pkg/datatype"
)

type (
	DetailWallet struct {
		Balance    float64             `json:"balance"`
		OwnedBy    string              `json:"owned_by"`
		ID         string              `json:"id"`
		EnabledAt  datatype.SqlIsoTime `json:"enabled_at,omitempty"`
		DisabledAt datatype.SqlIsoTime `json:"disabled_at,omitempty"`
		Status     enums.StatusWallet  `json:"status"`
	}
)

type (
	UpdateBalanceRequest struct {
		OwnedBy     string  `db:"owned_by" json:"owned_by"`
		Amount      float64 `db:"amount"`
		ReferenceID string  `db:"reference_id"`
	}
)
