package models

type Post struct {
	IDPost    int    `json:"id_post"`
	PostName  string `json:"post_name"`
	IsDeleted bool   `json:"is_deleted"`
}

type PostRequest struct {
	PostName string `json:"post_name"`
}
