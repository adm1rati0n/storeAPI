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

func GetAllAds(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `ad` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	//var users []models.User
	var ads []models.AdViewModel
	for rows.Next() {
		var ad models.Ad
		var adView models.AdViewModel
		err = rows.Scan(&ad.IDAd, &ad.AdPrice, &ad.AdDate, &ad.Agency, &ad.Employee, &ad.IsDeleted)
		if err != nil {
			panic(err)
		}
		adView.IDAd = ad.IDAd
		adView.AdPrice = ad.AdPrice
		adView.AdDate = ad.AdDate
		adView.IsDeleted = ad.IsDeleted
		adView.Agency = GetAgency(ad.Agency)
		adView.Employee = GetOneEmployee(ad.Employee)

		ads = append(ads, adView)
	}
	json.NewEncoder(w).Encode(ads)
}

func GetOneAd(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var ad models.Ad
	var adView models.AdViewModel

	if err := db.QueryRow("select * from `ad` where IsDeleted = 0 and ID_Ad = ?", id).Scan(&ad.IDAd, &ad.AdPrice, &ad.AdDate, &ad.Agency, &ad.Employee, &ad.IsDeleted); err != nil {
		panic(err)
	}
	adView.IDAd = ad.IDAd
	adView.AdPrice = ad.AdPrice
	adView.AdDate = ad.AdDate
	adView.IsDeleted = ad.IsDeleted
	adView.Agency = GetAgency(ad.Agency)
	adView.Employee = GetOneEmployee(ad.Employee)
	json.NewEncoder(w).Encode(adView)
}

func AddAd(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	var ad models.AdRequest
	err := json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		panic(err)
	}

	//Валидатор

	query := "call Ad_Insert(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, &ad.AdDate, &ad.AdPrice, &ad.Employee, &ad.Agency)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateAd(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	var ad models.AdRequest
	err = json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		panic(err)
	}

	//Валидатор

	query := "call Ad_Update(?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, &ad.AdDate, &ad.AdPrice, &ad.Employee, &ad.Agency, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteAd(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Ad_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
