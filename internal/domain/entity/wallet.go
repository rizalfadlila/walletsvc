package entity

import (
	"encoding/json"
	"github.com/julo/walletsvc/internal/domain/enums"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/julo/walletsvc/internal/pkg/datatype"
	uuid "github.com/satori/go.uuid"
)

type Wallet struct {
	Balance    float64             `db:"balance" json:"balance"`
	OwnedBy    string              `db:"owned_by" json:"owned_by"`
	ID         string              `db:"id" json:"id"`
	EnabledAt  datatype.SqlIsoTime `db:"enabled_at" json:"enabled_at"`
	DisabledAt datatype.SqlIsoTime `db:"disabled_at" json:"disabled_at"`
	Status     enums.StatusWallet  `db:"status" json:"status"`
}

func NewWallet(customerID string) Wallet {
	return Wallet{
		Balance: 0,
		OwnedBy: customerID,
		Status:  enums.Disabled,
		ID:      uuid.NewV4().String(),
	}
}

func (w *Wallet) ToDTODetailWallet() *dto.DetailWallet {
	return &dto.DetailWallet{
		Balance:    w.Balance,
		OwnedBy:    w.OwnedBy,
		ID:         w.ID,
		EnabledAt:  w.EnabledAt,
		DisabledAt: w.DisabledAt,
		Status:     w.Status,
	}
}

func (w *Wallet) Marshal() ([]byte, error) {
	return json.Marshal(*w)
}
