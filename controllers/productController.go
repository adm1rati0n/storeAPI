package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
	"strconv"
)

var db = dbConnection.DB

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Некорректный запрос")
	}
	rows, err := db.Query("select * from product where IsDeleted = 0 and ID_Product = $1", id)
	if err != nil {
		json.NewEncoder(w).Encode("Такого товара не существует")
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
