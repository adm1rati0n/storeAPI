package models

import "time"

type SupplyViewModel struct {
	IDSupply   int                 `json:"id_supply"`
	SupplyDate time.Time           `json:"supply_date"`
	Total      float32             `json:"total"`
	Employee   Employee            `json:"employee"`
	Supplier   Supplier            `json:"supplier"`
	Purchases  []PurchaseViewModel `json:"purchases"`
	IsDeleted  bool                `json:"is_deleted"`
}
