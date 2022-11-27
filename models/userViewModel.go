package models

type UserView struct {
	IDUser    int      `json:"id_user"`
	Login     string   `json:"login"`
	Password  string   `json:"password"`
	Role      Role     `json:"role"`
	Employee  Employee `json:"employee"`
	IsDeleted bool     `json:"is_deleted"`
}
