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

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from supplier where IsDeleted = 0")
	if err != nil {
		panic(err.Error())
	}
	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		err = rows.Scan(&supplier.IDSupplier, &supplier.SupplierName, &supplier.IsDeleted)
		if err != nil {
			panic(err.Error())
		}
		suppliers = append(suppliers, supplier)
	}
	json.NewEncoder(w).Encode(suppliers)
}

func GetOneSupplier(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Некорректный запрос")
	}

	var supplier models.Supplier

	if err := db.QueryRow("select * from `supplier` where IsDeleted = 0 and ID_Supplier = ?", id).Scan(&supplier.IDSupplier, &supplier.SupplierName, &supplier.IsDeleted); err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(supplier)
}

func AddSupplier(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	supplierName := r.FormValue("supplier_name")

	//Валидатор
	query := "call Supplier_Insert(?)"
	res, err := db.ExecContext(context.Background(), query, supplierName)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Некорректный запрос")
	}
	supplierName := r.FormValue("supplier_name")

	//Валидатор

	query := "call Supplier_Update(?,?)"
	res, err := db.ExecContext(context.Background(), query, id, supplierName)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Некорректный запрос")
	}
	query := "call Supplier_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}
