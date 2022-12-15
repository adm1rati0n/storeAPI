package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Cheque struct {
	IDCheque   int             `json:"id_cheque"`
	ChequeDate time.Time       `json:"cheque_date"`
	Total      decimal.Decimal `json:"total"`
	Employee   int             `json:"cheque_employee"`
	IsDeleted  bool            `json:"is_deleted"`
}
