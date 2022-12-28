package models

import (
	"github.com/shopspring/decimal"
)

type Cheque struct {
	IDCheque   int             `json:"id_cheque"`
	ChequeDate string          `json:"cheque_date"`
	Total      decimal.Decimal `json:"total"`
	Employee   int             `json:"cheque_employee"`
	IsDeleted  bool            `json:"is_deleted"`
}

type ChequeRequest struct {
	ChequeDate string `json:"cheque_date"`
	Employee   int    `json:"cheque_employee"`
}
