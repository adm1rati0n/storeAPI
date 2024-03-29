package models

import (
	"github.com/shopspring/decimal"
)

type SupplyViewModel struct {
	IDSupply    int                 `json:"id_supply"`
	SupplyDate  string              `json:"supply_date"`
	SupplyTotal decimal.Decimal     `json:"supply_total"`
	Employee    Employee            `json:"employee"`
	Supplier    Supplier            `json:"supplier"`
	Purchases   []PurchaseViewModel `json:"purchases"`
	IsDeleted   bool                `json:"is_deleted"`
}
