package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
	"strconv"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `product` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.IDProduct, &product.ProductName, &product.IsDeleted)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	product := GetProduct(id)

	json.NewEncoder(w).Encode(product)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}

	var product models.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	//Валидатор
	query := "call Product_Insert(?)"
	res, err := db.ExecContext(context.Background(), query, &product.ProductName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var product models.ProductRequest
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	query := "call Product_Update(?,?)"
	res, err := db.ExecContext(context.Background(), query, id, &product.ProductName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Product_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
