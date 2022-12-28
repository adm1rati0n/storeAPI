package models

import (
	"github.com/shopspring/decimal"
)

type ChequeViewModel struct {
	IDCheque   int             `json:"id_cheque"`
	ChequeDate string          `json:"cheque_date"`
	Total      decimal.Decimal `json:"total"`
	Employee   Employee        `json:"cheque_employee"`
	Products   []SaleViewModel `json:"products"`
	IsDeleted  bool            `json:"is_deleted"`
}
