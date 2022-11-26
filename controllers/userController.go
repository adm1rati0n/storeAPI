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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `user` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.IDUser, &user.Login, &user.Password, &user.UserEmployee, &user.UserRole, &user.IsDeleted)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var user models.User

	if err := db.QueryRow("select * from `user` where IsDeleted = 0 and ID_Product = ?", id).Scan(&user.IDUser, &user.Login, &user.Password, &user.UserEmployee, &user.UserRole, &user.IsDeleted); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	login := r.FormValue("login")
	password := r.FormValue("password")
	employeeID := r.FormValue("user_employee")
	roleID := r.FormValue("user_role")

	//Валидатор

	password = HashPassword(password)
	query := "call User_Insert(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, login, password, employeeID, roleID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	login := r.FormValue("login")
	password := r.FormValue("password")
	employeeID := r.FormValue("user_employee")
	roleID := r.FormValue("user_role")

	//Валидатор

	password = HashPassword(password)

	query := "call User_Update(?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, id, login, password, employeeID, roleID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call User_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
