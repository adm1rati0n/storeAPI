package models

type Supplier struct {
	IDSupplier   int    `json:"id_supplier"`
	SupplierName string `json:"supplier_name"`
	IsDeleted    bool   `json:"is_deleted"`
}

type SupplierRequest struct {
	SupplierName string `json:"supplier_name"`
}
