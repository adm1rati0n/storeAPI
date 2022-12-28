package models

type EmployeeRequest struct {
	EmployeeSurname        string `json:"employee_surname"`
	EmployeeName           string `json:"employee_name"`
	EmployeeMiddlename     string `json:"employee_middlename"`
	EmployeePassportSeries string `json:"employee_passport_series"`
	EmployeePassportNumber string `json:"employee_passport_number"`
	Post                   int    `json:"employee_post"`
}
