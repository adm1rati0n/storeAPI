package models

type Ad struct {
	IDAd      int     `json:"id_ad"`
	AdDate    string  `json:"ad_date"`
	AdPrice   float32 `json:"ad_price"`
	Agency    int     `json:"agency"`
	Employee  int     `json:"employee"`
	IsDeleted bool    `json:"is_deleted"`
}

type AdRequest struct {
	AdDate   string  `json:"ad_date"`
	AdPrice  float32 `json:"ad_price"`
	Agency   int     `json:"agency"`
	Employee int     `json:"employee"`
}
