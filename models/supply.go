package models

import (
	"github.com/shopspring/decimal"
)

type Supply struct {
	IDSupply    int             `json:"id_supply"`
	SupplyDate  string          `json:"supply_date"`
	SupplyTotal decimal.Decimal `json:"supply_total"`
	Employee    int             `json:"employee"`
	Supplier    int             `json:"supplier"`
	IsDeleted   bool            `json:"is_deleted"`
}

type SupplyRequest struct {
	SupplyDate string `json:"supply_date"`
	Employee   int    `json:"employee"`
	Supplier   int    `json:"supplier"`
}
