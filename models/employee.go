package models

type Employee struct {
	IDEmployee             int    `json:"id_employee"`
	EmployeeSurname        string `json:"employee_surname"`
	EmployeeName           string `json:"employee_name"`
	EmployeeMiddlename     string `json:"employee_middlename"`
	EmployeePassportSeries string `json:"employee_passport_series"`
	EmployeePassportNumber string `json:"employee_passport_number"`
	EmployeePost           int    `json:"employee_post"`
	IsDeleted              bool   `json:"is_deleted"`
}
