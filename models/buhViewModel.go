package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type BuhViewModel struct {
	IDBuh        int             `json:"id_buh"`
	StartingDate time.Time       `json:"starting_date"`
	EndingDate   time.Time       `json:"ending_date"`
	Earnings     decimal.Decimal `json:"earnings"`
	Expenses     decimal.Decimal `json:"expenses"`
	Taxes        decimal.Decimal `json:"taxes"`
	Profit       decimal.Decimal `json:"profit"`
	Employee     Employee        `json:"employee"`
	IsDeleted    bool            `json:"is_deleted"`
}
