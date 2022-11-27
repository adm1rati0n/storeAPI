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

func GetAllAgencies(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `agency` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var agencies []models.Agency
	for rows.Next() {
		var agency models.Agency
		err = rows.Scan(&agency.IDAgency, &agency.AgencyName, &agency.IsDeleted)
		if err != nil {
			panic(err)
		}
		agencies = append(agencies, agency)
	}
	json.NewEncoder(w).Encode(agencies)
}

func GetOneAgency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	agency := GetAgency(id)
	json.NewEncoder(w).Encode(agency)
}

func AddAgency(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	agencyName := r.FormValue("agency_name")

	//Валидатор
	query := "call Agency_Insert(?)"
	res, err := db.ExecContext(context.Background(), query, agencyName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateAgency(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	agencyName := r.FormValue("agency_name")

	//Валидатор

	query := "call Agency_Update(?,?)"
	res, err := db.ExecContext(context.Background(), query, id, agencyName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteAgency(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Agency_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
