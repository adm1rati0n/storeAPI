package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
	"strconv"
)

func GetFilteredEmployees(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select * from `employee` where IsDeleted = 0 and Post_ID = ?", id)
	if err != nil {
		panic(err)
	}
	var employees []models.EmployeeViewModel
	for rows.Next() {
		var employee models.Employee
		var employeeView models.EmployeeViewModel
		err = rows.Scan(&employee.IDEmployee, &employee.EmployeeSurname, &employee.EmployeeName,
			&employee.EmployeeMiddlename, &employee.EmployeePassportSeries, &employee.EmployeePassportNumber, &employee.Post, &employee.IsDeleted)
		if err != nil {
			panic(err)
		}

		employeeView.IDEmployee = employee.IDEmployee
		employeeView.EmployeeSurname = employee.EmployeeSurname
		employeeView.EmployeeName = employee.EmployeeName
		employeeView.EmployeeMiddlename = employee.EmployeeMiddlename
		employeeView.EmployeePassportSeries = employee.EmployeePassportSeries
		employeeView.EmployeePassportNumber = employee.EmployeePassportNumber
		employeeView.IsDeleted = employee.IsDeleted
		employeeView.Post = GetPost(employee.Post)

		employees = append(employees, employeeView)
	}
	json.NewEncoder(w).Encode(employees)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `employee` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}
	var employees []models.EmployeeViewModel
	for rows.Next() {
		var employee models.Employee
		var employeeView models.EmployeeViewModel
		err = rows.Scan(&employee.IDEmployee, &employee.EmployeeSurname, &employee.EmployeeName,
			&employee.EmployeeMiddlename, &employee.EmployeePassportSeries, &employee.EmployeePassportNumber, &employee.Post, &employee.IsDeleted)
		if err != nil {
			panic(err)
		}

		employeeView.IDEmployee = employee.IDEmployee
		employeeView.EmployeeSurname = employee.EmployeeSurname
		employeeView.EmployeeName = employee.EmployeeName
		employeeView.EmployeeMiddlename = employee.EmployeeMiddlename
		employeeView.EmployeePassportSeries = employee.EmployeePassportSeries
		employeeView.EmployeePassportNumber = employee.EmployeePassportNumber
		employeeView.IsDeleted = employee.IsDeleted
		employeeView.Post = GetPost(employee.Post)

		employees = append(employees, employeeView)
	}
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var employee models.Employee
	var employeeView models.EmployeeViewModel

	if err := db.QueryRow("select * from `employee` where IsDeleted = 0 and ID_Employee = ?", id).Scan(
		&employee.IDEmployee, &employee.EmployeeSurname, &employee.EmployeeName,
		&employee.EmployeeMiddlename, &employee.EmployeePassportSeries, &employee.EmployeePassportNumber,
		&employee.Post, &employee.IsDeleted); err != nil {
		panic(err)
	}

	employeeView.IDEmployee = employee.IDEmployee
	employeeView.EmployeeSurname = employee.EmployeeSurname
	employeeView.EmployeeName = employee.EmployeeName
	employeeView.EmployeeMiddlename = employee.EmployeeMiddlename
	employeeView.EmployeePassportSeries = employee.EmployeePassportSeries
	employeeView.EmployeePassportNumber = employee.EmployeePassportNumber
	employeeView.IsDeleted = employee.IsDeleted
	employeeView.Post = GetPost(employee.Post)
	json.NewEncoder(w).Encode(employeeView)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Метод вызван")
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	fmt.Println(r.Body)

	var employee models.EmployeeRequest

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		panic(err)
	}

	query := "call Employee_Insert(?,?,?,?,?,?)"
	res, err := db.ExecContext(
		context.Background(), query, &employee.EmployeeSurname, &employee.EmployeeName,
		&employee.EmployeeMiddlename, &employee.EmployeePassportSeries, &employee.EmployeePassportNumber,
		&employee.Post)

	if err != nil {
		panic(err)
	}
	fmt.Println("Запись сохранена")
	json.NewEncoder(w).Encode(res)
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	var employee models.EmployeeRequest

	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		panic(err)
	}

	query := "call Employee_Update(?,?,?,?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, &employee.EmployeeSurname, &employee.EmployeeName,
		&employee.EmployeeMiddlename, &employee.EmployeePassportSeries, &employee.EmployeePassportNumber,
		&employee.Post, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Employee_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
