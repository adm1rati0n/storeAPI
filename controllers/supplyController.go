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

func GetAllSupplies(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `supply` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var supplies []models.SupplyViewModel
	for rows.Next() {
		var supply models.Supply
		var supplyView models.SupplyViewModel
		err = rows.Scan(&supply.IDSupply, &supply.SupplyDate, &supply.SupplyTotal,
			&supply.Employee, &supply.Supplier, &supply.IsDeleted)
		if err != nil {
			panic(err)
		}

		supplyView.IDSupply = supply.IDSupply
		supplyView.SupplyDate = supply.SupplyDate
		supplyView.Total = supply.SupplyTotal
		supplyView.IsDeleted = supply.IsDeleted
		supplyView.Supplier = GetSupplier(supply.Supplier)
		supplyView.Employee = GetOneEmployee(supply.Employee)
		supplyView.Purchases = GetPurchases(supply.IDSupply)

		supplies = append(supplies, supplyView)
	}
	json.NewEncoder(w).Encode(supplies)
}

func GetOneSupply(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var supply models.Supply
	var supplyView models.SupplyViewModel

	if err := db.QueryRow("select * from `supply` where IsDeleted = 0 and ID_Supply = ?", id).Scan(
		&supply.IDSupply, &supply.SupplyDate, &supply.SupplyTotal,
		&supply.Employee, &supply.Supplier, &supply.IsDeleted); err != nil {
		panic(err)
	}

	supplyView.IDSupply = supply.IDSupply
	supplyView.SupplyDate = supply.SupplyDate
	supplyView.Total = supply.SupplyTotal
	supplyView.IsDeleted = supply.IsDeleted
	supplyView.Supplier = GetSupplier(supply.Supplier)
	supplyView.Employee = GetOneEmployee(supply.Employee)
	supplyView.Purchases = GetPurchases(supply.IDSupply)
	json.NewEncoder(w).Encode(supplyView)
}

func AddSupply(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	date := r.FormValue("supply_date")
	employeeID := r.FormValue("employee")
	supplierID := r.FormValue("supplier")

	//Валидатор

	query := "call Supply_Insert(?,?,?)"
	res, err := db.ExecContext(context.Background(), query, date, employeeID, supplierID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteSupply(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Supply_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
