package models

type Sale struct {
	IDSale    int     `json:"id_sale"`
	Amount    int     `json:"amount"`
	Price     float32 `json:"price"`
	Product   int     `json:"product"`
	Cheque    int     `json:"cheque"`
	IsDeleted bool    `json:"is_deleted"`
}
