package models

type Role struct {
	IDRole    int    `json:"id_role"`
	RoleName  string `json:"role_name"`
	IsDeleted bool   `json:"is_deleted"`
}
