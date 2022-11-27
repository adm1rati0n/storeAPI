package models

import "time"

type AdViewModel struct {
	IDAd      int       `json:"id_ad"`
	AdDate    time.Time `json:"ad_date"`
	AdPrice   float32   `json:"ad_price"`
	Agency    Agency    `json:"agency"`
	Employee  Employee  `json:"employee"`
	IsDeleted bool      `json:"is_deleted"`
}
