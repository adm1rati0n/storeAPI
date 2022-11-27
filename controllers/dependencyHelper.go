package controllers

import (
	"storeAPI/dbConnection"
	"storeAPI/models"
)

func GetOneRole(id int) models.Role {
	db := dbConnection.DB
	var role models.Role
	if err := db.QueryRow("select * from `role` where IsDeleted = 0 and ID_Role = ?", id).Scan(&role.IDRole, &role.RoleName, &role.IsDeleted); err != nil {
		panic(err)
	}
	return role
}

func GetOneEmployee(id int) models.Employee {
	db := dbConnection.DB
	var employee models.Employee
	if err := db.QueryRow("select * from `employee` where IsDeleted = 0 and ID_Employee = ?", id).Scan(
		&employee.IDEmployee, &employee.EmployeeSurname, &employee.EmployeeName, &employee.EmployeeMiddlename,
		&employee.EmployeePassportSeries, &employee.EmployeePassportNumber, &employee.Post, &employee.IsDeleted); err != nil {
		panic(err)
	}
	return employee
}

func GetAgency(id int) models.Agency {
	db := dbConnection.DB

	var agency models.Agency

	if err := db.QueryRow("select * from `agency` where IsDeleted = 0 and ID_Agency = ?", id).Scan(&agency.IDAgency, &agency.AgencyName, &agency.IsDeleted); err != nil {
		panic(err)
	}
	return agency
}
