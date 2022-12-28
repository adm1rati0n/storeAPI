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
	var users []models.UserView
	for rows.Next() {
		var user models.User
		var userView models.UserView
		err = rows.Scan(&user.IDUser, &user.Login, &user.Password, &user.Employee, &user.Role, &user.IsDeleted)
		if err != nil {
			panic(err)
		}
		userView.IDUser = user.IDUser
		userView.Login = user.Login
		userView.Password = user.Password
		userView.IsDeleted = user.IsDeleted
		userView.Role = GetOneRole(user.Role)
		userView.Employee = GetOneEmployee(user.Employee)

		users = append(users, userView)
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
	var userView models.UserView

	if err := db.QueryRow("select * from `user` where IsDeleted = 0 and ID_User = ?", id).Scan(&user.IDUser, &user.Login, &user.Password, &user.Employee, &user.Role, &user.IsDeleted); err != nil {
		panic(err)
	}
	userView.IDUser = user.IDUser
	userView.Login = user.Login
	userView.Password = user.Password
	userView.IsDeleted = user.IsDeleted
	userView.Role = GetOneRole(user.Role)
	userView.Employee = GetOneEmployee(user.Employee)
	json.NewEncoder(w).Encode(userView)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}

	var user models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	var password = HashPassword(user.Password)
	query := "call User_Insert(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, &user.Login, password, &user.Employee, &user.Role)
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
	var user models.UserRequest
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	//Валидатор

	var password = HashPassword(user.Password)

	query := "call User_Update(?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, &user.Login, password, &user.Employee, &user.Role, id)
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

func GetRoles(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `role` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err = rows.Scan(&role.IDRole, &role.RoleName, &role.IsDeleted)
		if err != nil {
			panic(err)
		}
		roles = append(roles, role)
	}
	json.NewEncoder(w).Encode(roles)
}
