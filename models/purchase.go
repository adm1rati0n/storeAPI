package models

import "github.com/shopspring/decimal"

type Purchase struct {
	IDPurchase     int             `json:"id_purchase"`
	PurchaseAmount int             `json:"purchase_amount"`
	PurchasePrice  decimal.Decimal `json:"purchase_price"`
	Product        int             `json:"product"`
	Supply         int             `json:"supply"`
	IsDeleted      bool            `json:"is_deleted"`
}
