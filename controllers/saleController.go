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

func GetAllSales(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `sale` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var sales []models.SaleViewModel
	for rows.Next() {
		var sale models.Sale
		var saleView models.SaleViewModel
		err = rows.Scan(&sale.IDSale, &sale.Amount, &sale.Price,
			&sale.Product, &sale.Cheque, &sale.IsDeleted)
		if err != nil {
			panic(err)
		}

		saleView.IDSale = sale.IDSale
		saleView.Amount = sale.Amount
		saleView.Price = sale.Price
		saleView.IsDeleted = sale.IsDeleted
		saleView.Cheque = sale.Cheque
		saleView.Product = GetProduct(sale.Product)

		sales = append(sales, saleView)
	}
	json.NewEncoder(w).Encode(sales)
}

func GetOneSale(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var sale models.Sale
	var saleView models.SaleViewModel

	if err := db.QueryRow("select * from `sale` where IsDeleted = 0 and ID_Sale = ?", id).Scan(
		&sale.IDSale, &sale.Amount, &sale.Price,
		&sale.Product, &sale.Cheque, &sale.IsDeleted); err != nil {
		panic(err)
	}

	saleView.IDSale = sale.IDSale
	saleView.Amount = sale.Amount
	saleView.Price = sale.Price
	saleView.IsDeleted = sale.IsDeleted
	saleView.Cheque = sale.Cheque
	saleView.Product = GetProduct(sale.Product)
	json.NewEncoder(w).Encode(saleView)
}

func AddSale(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	amount := r.FormValue("amount")
	price := r.FormValue("price")
	productID := r.FormValue("product")
	supplyID := r.FormValue("cheque")

	//Валидатор

	query := "call Sale_Insert(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, amount, price, productID, supplyID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateSale(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	amount := r.FormValue("amount")
	price := r.FormValue("price")
	productID := r.FormValue("product")
	supplyID := r.FormValue("cheque")

	//Валидатор

	query := "call Sale_Update(?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, id, amount, price, productID, supplyID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteSale(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Sale_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
