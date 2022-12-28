package models

type Product struct {
	IDProduct   int    `json:"id_product"`
	ProductName string `json:"product_name"`
	IsDeleted   bool   `json:"is_deleted"`
}

type ProductRequest struct {
	ProductName string `json:"product_name"`
}
