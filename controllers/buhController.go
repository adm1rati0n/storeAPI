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

func GetAllBuh(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `buh` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var buhs []models.BuhViewModel
	for rows.Next() {
		var buh models.Buh
		var buhView models.BuhViewModel
		err = rows.Scan(&buh.IDBuh, &buh.StartingDate, &buh.EndingDate, &buh.Earnings, &buh.Expenses, &buh.Taxes, &buh.Profit, &buh.Employee, &buh.IsDeleted)
		if err != nil {
			panic(err)
		}
		employee := GetOneEmployee(buh.Employee)
		buhView.IDBuh = buh.IDBuh
		buhView.StartingDate = buh.StartingDate
		buhView.EndingDate = buh.EndingDate
		buhView.Earnings = buh.Earnings
		buhView.Expenses = buh.Expenses
		buhView.Taxes = buh.Taxes
		buhView.Profit = buh.Profit
		buhView.IsDeleted = buh.IsDeleted
		buhView.Employee = employee

		buhs = append(buhs, buhView)
	}
	json.NewEncoder(w).Encode(buhs)
}

func GetOneBuh(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var buh models.Buh
	var buhView models.BuhViewModel

	if err := db.QueryRow("select * from `buh` where IsDeleted = 0 and ID_Buh = ?", id).Scan(
		&buh.IDBuh, &buh.StartingDate, &buh.EndingDate, &buh.Earnings, &buh.Expenses, &buh.Taxes, &buh.Profit,
		&buh.Employee, &buh.IsDeleted); err != nil {
		panic(err)
	}
	employee := GetOneEmployee(buh.Employee)
	buhView.IDBuh = buh.IDBuh
	buhView.StartingDate = buh.StartingDate
	buhView.EndingDate = buh.EndingDate
	buhView.Earnings = buh.Earnings
	buhView.Expenses = buh.Expenses
	buhView.Taxes = buh.Taxes
	buhView.Profit = buh.Profit
	buhView.IsDeleted = buh.IsDeleted
	buhView.Employee = employee
	json.NewEncoder(w).Encode(buhView)
}

func AddBuh(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	startingDate := r.FormValue("starting_date")
	endingDate := r.FormValue("ending_date")
	employeeID := r.FormValue("employee")

	//Валидатор

	query := "call Buh_Insert(?,?,?)"
	res, err := db.ExecContext(context.Background(), query, startingDate, endingDate, employeeID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateBuh(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	startingDate := r.FormValue("starting_date")
	endingDate := r.FormValue("ending_date")
	employeeID := r.FormValue("employee")

	//Валидатор

	query := "call Buh_Update(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, id, startingDate, endingDate, employeeID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteBuh(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Buh_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
