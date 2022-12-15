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

func GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `purchase` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var purchases []models.PurchaseViewModel
	for rows.Next() {
		var purchase models.Purchase
		var purchaseView models.PurchaseViewModel
		err = rows.Scan(&purchase.IDPurchase, &purchase.PurchaseAmount, &purchase.PurchasePrice,
			&purchase.Product, &purchase.Supply, &purchase.IsDeleted)
		if err != nil {
			panic(err)
		}

		purchaseView.IDPurchase = purchase.IDPurchase
		purchaseView.PurchaseAmount = purchase.PurchaseAmount
		purchaseView.PurchasePrice = purchase.PurchasePrice
		purchaseView.IsDeleted = purchase.IsDeleted
		purchaseView.Supply = purchase.Supply
		purchaseView.Product = GetProduct(purchase.Product)

		purchases = append(purchases, purchaseView)
	}
	json.NewEncoder(w).Encode(purchases)
}

func GetOnePurchase(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var purchase models.Purchase
	var purchaseView models.PurchaseViewModel

	if err := db.QueryRow("select * from `purchase` where IsDeleted = 0 and ID_Purchase = ?", id).Scan(
		&purchase.IDPurchase, &purchase.PurchaseAmount, &purchase.PurchasePrice,
		&purchase.Product, &purchase.Supply, &purchase.IsDeleted); err != nil {
		panic(err)
	}

	purchaseView.IDPurchase = purchase.IDPurchase
	purchaseView.PurchaseAmount = purchase.PurchaseAmount
	purchaseView.PurchasePrice = purchase.PurchasePrice
	purchaseView.IsDeleted = purchase.IsDeleted
	purchaseView.Supply = purchase.Supply
	purchaseView.Product = GetProduct(purchase.Product)
	json.NewEncoder(w).Encode(purchaseView)
}

func AddPurchase(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	amount := r.FormValue("purchase_amount")
	price := r.FormValue("purchase_price")
	productID := r.FormValue("product")
	supplyID := r.FormValue("supply")

	//Валидатор

	query := "call Purchase_Insert(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, amount, price, productID, supplyID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	amount := r.FormValue("purchase_amount")
	price := r.FormValue("purchase_price")
	productID := r.FormValue("product")
	supplyID := r.FormValue("supply")

	//Валидатор

	query := "call Purchase_Update(?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, amount, price, productID, supplyID, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeletePurchase(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Purchase_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
