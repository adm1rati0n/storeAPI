package models

import "time"

type Supply struct {
	IDSupply    int       `json:"id_supply"`
	SupplyDate  time.Time `json:"supply_date"`
	SupplyTotal float32   `json:"supply_total"`
	Employee    int       `json:"employee"`
	Supplier    int       `json:"supplier"`
	IsDeleted   bool      `json:"is_deleted"`
}
