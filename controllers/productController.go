package controllers

import (
	"encoding/json"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from product where IsDeleted = 0")
	if err != nil {
		panic(err.Error())
	}
	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.IDProduct, &product.ProductName, &product.IsDeleted)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {

}
