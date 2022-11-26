package controllers

import (
	"encoding/json"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `post` where IsDeleted = 0")
	if err != nil {
		panic(err.Error())
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.IDPost, &post.PostName, &post.IsDeleted)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
