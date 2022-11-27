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

func GetAllCheques(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `cheque` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var cheques []models.ChequeViewModel
	for rows.Next() {
		var cheque models.Cheque
		var chequeView models.ChequeViewModel
		err = rows.Scan(&cheque.IDCheque, &cheque.ChequeDate, &cheque.Total,
			&cheque.Employee, &cheque.IsDeleted)
		if err != nil {
			panic(err)
		}

		chequeView.IDCheque = cheque.IDCheque
		chequeView.ChequeDate = cheque.ChequeDate
		chequeView.Total = cheque.Total
		chequeView.IsDeleted = cheque.IsDeleted
		chequeView.Employee = GetOneEmployee(cheque.Employee)
		chequeView.Products = GetSales(cheque.IDCheque)

		cheques = append(cheques, chequeView)
	}
	json.NewEncoder(w).Encode(cheques)
}

func GetOneCheque(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var cheque models.Cheque
	var chequeView models.ChequeViewModel

	if err := db.QueryRow("select * from `cheque` where IsDeleted = 0 and ID_Cheque = ?", id).Scan(
		&cheque.IDCheque, &cheque.ChequeDate, &cheque.Total,
		&cheque.Employee, &cheque.IsDeleted); err != nil {
		panic(err)
	}

	chequeView.IDCheque = cheque.IDCheque
	chequeView.ChequeDate = cheque.ChequeDate
	chequeView.Total = cheque.Total
	chequeView.IsDeleted = cheque.IsDeleted
	chequeView.Employee = GetOneEmployee(cheque.Employee)
	chequeView.Products = GetSales(cheque.IDCheque)
	json.NewEncoder(w).Encode(chequeView)
}

func AddCheque(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	date := r.FormValue("cheque_date")
	employeeID := r.FormValue("cheque_employee")

	//Валидатор

	query := "call Cheque_Insert(?,?)"
	res, err := db.ExecContext(context.Background(), query, date, employeeID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteCheque(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Cheque_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
