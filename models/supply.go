package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Supply struct {
	IDSupply    int             `json:"id_supply"`
	SupplyDate  time.Time       `json:"supply_date"`
	SupplyTotal decimal.Decimal `json:"supply_total"`
	Employee    int             `json:"employee"`
	Supplier    int             `json:"supplier"`
	IsDeleted   bool            `json:"is_deleted"`
}
