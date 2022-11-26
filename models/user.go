package models

type User struct {
	IDUser       int    `json:"id_user"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	UserRole     int    `json:"user_role"`
	UserEmployee int    `json:"user_employee"`
	IsDeleted    bool   `json:"is_deleted"`
}
