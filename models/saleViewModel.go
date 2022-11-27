package models

type SaleViewModel struct {
	IDSale    int     `json:"id_sale"`
	Amount    int     `json:"amount"`
	Price     float32 `json:"price"`
	Product   Product `json:"product"`
	Cheque    int     `json:"cheque"`
	IsDeleted bool    `json:"is_deleted"`
}
