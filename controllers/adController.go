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
		agency := GetAgency(ad.Agency)
		employee := GetOneEmployee(ad.Employee)
		adView.IDAd = ad.IDAd
		adView.AdPrice = ad.AdPrice
		adView.AdDate = ad.AdDate
		adView.IsDeleted = ad.IsDeleted
		adView.Agency = agency
		adView.Employee = employee

		//users = append(users, userView)

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
	agency := GetAgency(ad.Agency)
	employee := GetOneEmployee(ad.Employee)
	adView.IDAd = ad.IDAd
	adView.AdPrice = ad.AdPrice
	adView.AdDate = ad.AdDate
	adView.IsDeleted = ad.IsDeleted
	adView.Agency = agency
	adView.Employee = employee
	json.NewEncoder(w).Encode(adView)
}

func AddAd(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	date := r.FormValue("ad_date")
	price := r.FormValue("ad_price")
	employeeID := r.FormValue("employee")
	agencyID := r.FormValue("agency")

	//Валидатор

	query := "call Ad_Insert(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, date, price, employeeID, agencyID)
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
	date := r.FormValue("ad_date")
	price := r.FormValue("ad_price")
	employeeID := r.FormValue("employee")
	agencyID := r.FormValue("agency")

	//Валидатор

	query := "call Ad_Update(?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, id, date, price, employeeID, agencyID)
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
