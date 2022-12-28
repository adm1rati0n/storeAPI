package models

import (
	"github.com/shopspring/decimal"
)

type BuhViewModel struct {
	IDBuh        int             `json:"id_buh"`
	StartingDate string          `json:"starting_date"`
	EndingDate   string          `json:"ending_date"`
	Earnings     decimal.Decimal `json:"earnings"`
	Expenses     decimal.Decimal `json:"expenses"`
	Taxes        decimal.Decimal `json:"taxes"`
	Profit       decimal.Decimal `json:"profit"`
	Employee     Employee        `json:"employee"`
	IsDeleted    bool            `json:"is_deleted"`
}
