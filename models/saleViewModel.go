package models

import "github.com/shopspring/decimal"

type SaleViewModel struct {
	IDSale    int             `json:"id_sale"`
	Amount    int             `json:"amount"`
	Price     decimal.Decimal `json:"price"`
	Product   Product         `json:"product"`
	Cheque    int             `json:"cheque"`
	IsDeleted bool            `json:"is_deleted"`
}
