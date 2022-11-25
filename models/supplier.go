package models

type Supplier struct {
	IDSupplier   int    `json:"id_supplier"`
	SupplierName string `json:"supplier_name"`
	IsDeleted    string `json:"is_deleted"`
}
