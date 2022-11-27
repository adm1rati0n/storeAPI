package models

import "time"

type Ad struct {
	IDAd      int       `json:"id_ad"`
	AdDate    time.Time `json:"ad_date"`
	AdPrice   float32   `json:"ad_price"`
	Agency    int       `json:"agency"`
	Employee  int       `json:"employee"`
	IsDeleted bool      `json:"is_deleted"`
}
