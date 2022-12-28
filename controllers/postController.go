package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"storeAPI/dbConnection"
	"storeAPI/models"
	"strconv"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	rows, err := db.Query("select * from `post` where IsDeleted = 0")
	if err != nil {
		panic(err)
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.IDPost, &post.PostName, &post.IsDeleted)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func GetOnePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	post := GetPost(id)
	json.NewEncoder(w).Encode(post)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}

	var post models.PostRequest
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		panic(err)
	}

	query := "call Post_Insert(?)"
	res, err := db.ExecContext(context.Background(), query, &post.PostName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	if r.Body == nil {
		json.NewEncoder(w).Encode("Поля ввода не заполнены")
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var post models.PostRequest
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		panic(err)
	}

	query := "call Post_Update(?,?)"
	res, err := db.ExecContext(context.Background(), query, id, &post.PostName)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := dbConnection.DB
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	query := "call Post_Delete(?)"
	res, err := db.ExecContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(res)
}
