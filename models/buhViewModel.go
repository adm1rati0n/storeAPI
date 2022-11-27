package models

import "time"

type BuhViewModel struct {
	IDBuh        int       `json:"id_buh"`
	StartingDate time.Time `json:"starting_date"`
	EndingDate   time.Time `json:"ending_date"`
	Earnings     float32   `json:"earnings"`
	Expenses     float32   `json:"expenses"`
	Taxes        float32   `json:"taxes"`
	Profit       float32   `json:"profit"`
	Employee     Employee  `json:"employee"`
	IsDeleted    bool      `json:"is_deleted"`
}
