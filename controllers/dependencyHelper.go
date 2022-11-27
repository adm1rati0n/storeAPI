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

func GetPost(id int) models.Post {
	db := dbConnection.DB
	var post models.Post
	if err := db.QueryRow("select * from `post` where IsDeleted = 0 and ID_Post = ?", id).Scan(&post.IDPost, &post.PostName, &post.IsDeleted); err != nil {
		panic(err)
	}
	return post
}

func GetProduct(id int) models.Product {
	db := dbConnection.DB
	var product models.Product

	if err := db.QueryRow("select * from `product` where IsDeleted = 0 and ID_Product = ?", id).Scan(&product.IDProduct, &product.ProductName, &product.IsDeleted); err != nil {
		panic(err)
	}
	return product
}

func GetSales(id int) []models.SaleViewModel {
	db := dbConnection.DB
	rows, err := db.Query("select * from `sale` where IsDeleted = 0 and Cheque_ID = ?", id)
	if err != nil {
		panic(err)
	}
	var sales []models.SaleViewModel
	for rows.Next() {
		var sale models.Sale
		var saleView models.SaleViewModel
		err = rows.Scan(&sale.IDSale, &sale.Amount, &sale.Price,
			&sale.Product, &sale.Cheque, &sale.IsDeleted)
		if err != nil {
			panic(err)
		}

		saleView.IDSale = sale.IDSale
		saleView.Amount = sale.Amount
		saleView.Price = sale.Price
		saleView.IsDeleted = sale.IsDeleted
		saleView.Cheque = sale.Cheque
		saleView.Product = GetProduct(sale.Product)

		sales = append(sales, saleView)
	}
	return sales
}
