package main

import (
	"net/http"
	"storeAPI/controllers"
	"storeAPI/dbConnection"
)

func main() {
	dbConnection.DBConnect()
	http.Handle("/api", controllers.ValidateJWT(controllers.GetPosts))
	http.HandleFunc("/jwt", controllers.GetJWT)
	http.ListenAndServe("localhost:4000", nil)
}
