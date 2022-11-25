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

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	productName := r.FormValue("product_name")

	//Валидатор

	err := db.QueryRow("CALL Product_Insert($1)", productName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("Товар добавлен")
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Некорректный запрос")
	}
	productName := r.FormValue("product_name")

	//Валидатор

	e := db.QueryRow("CALL Product_Update($1,$2)", id, productName)
	if e != nil {
		panic(e)
	}
	json.NewEncoder(w).Encode("Товар изменен")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Некорректный запрос")
	}
	e := db.QueryRow("CALL Product_Delete($1)", id)
	if e != nil {
		panic(e)
	}
	json.NewEncoder(w).Encode("Товар удален")
}
