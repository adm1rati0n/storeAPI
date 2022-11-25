package models

import "time"

type Cheque struct {
	IDCheque       int       `json:"id_cheque"`
	ChequeDate     time.Time `json:"cheque_date"`
	Total          float32   `json:"total"`
	ChequeEmployee Employee  `json:"cheque_employee"`
	IsDeleted      bool      `json:"is_deleted"`
}
