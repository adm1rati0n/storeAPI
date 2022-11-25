package models

type Agency struct {
	IDAgency   int    `json:"id_agency"`
	AgencyName string `json:"agency_name"`
	IsDeleted  bool   `json:"is_deleted"`
}
