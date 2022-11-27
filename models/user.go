package models

type User struct {
	IDUser    int    `json:"id_user"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	Employee  int    `json:"employee"`
	IsDeleted bool   `json:"is_deleted"`
}
