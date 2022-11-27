package models

import "time"

type ChequeViewModel struct {
	IDCheque   int             `json:"id_cheque"`
	ChequeDate time.Time       `json:"cheque_date"`
	Total      float32         `json:"total"`
	Employee   Employee        `json:"cheque_employee"`
	Products   []SaleViewModel `json:"products"`
	IsDeleted  bool            `json:"is_deleted"`
}
