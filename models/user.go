package models

type User struct {
	IDUser       int      `json:"id_user"`
	Login        string   `json:"login"`
	Password     string   `json:"password"`
	UserRole     Role     `json:"user_role"`
	UserEmployee Employee `json:"user_employee"`
}
