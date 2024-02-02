package entity

import (
	"github.com/julo/walletsvc/internal/pkg/datatype"
)

type Transaction struct {
	Amount       float64             `db:"amount"`
	TransactedAt datatype.SqlIsoTime `db:"transacted_at"`
	ReferenceID  string              `db:"reference_id"`
	ID           string              `db:"id"`
	Type         string              `db:"type"`
	Status       string              `db:"status"`
}
