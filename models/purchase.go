package models

type Purchase struct {
	IDPurchase     int     `json:"id_purchase"`
	PurchaseAmount int     `json:"purchase_amount"`
	PurchasePrice  float32 `json:"purchase_price"`
	Product        int     `json:"product"`
	Supply         int     `json:"supply"`
	IsDeleted      bool    `json:"is_deleted"`
}
