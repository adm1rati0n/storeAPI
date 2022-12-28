package models

import "github.com/shopspring/decimal"

type Sale struct {
	IDSale    int             `json:"id_sale"`
	Amount    int             `json:"amount"`
	Price     decimal.Decimal `json:"price"`
	Product   int             `json:"product"`
	Cheque    int             `json:"cheque"`
	IsDeleted bool            `json:"is_deleted"`
}

type SaleRequest struct {
	Amount  int             `json:"amount"`
	Price   decimal.Decimal `json:"price"`
	Product int             `json:"product"`
	Cheque  int             `json:"cheque"`
}
