package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"storeAPI/controllers"
	"storeAPI/dbConnection"
)

func main() {
	dbConnection.DBConnect()
	r := mux.NewRouter()
	r.Handle("/api", controllers.ValidateJWT(controllers.GetPosts))
	r.HandleFunc("/jwt", controllers.GetJWT)
	r.HandleFunc("/posts", controllers.GetPosts)

	//Товары
	r.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/add", controllers.AddProduct).Methods("POST")
	r.HandleFunc("/products/edit/{id}", controllers.UpdateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.GetOneProduct).Methods("GET")
	r.HandleFunc("/products/delete/{id}", controllers.DeleteProduct).Methods("GET")

	//Рекламные агенства
	r.HandleFunc("/agencies", controllers.GetAllAgencies).Methods("GET")
	r.HandleFunc("/agencies/add", controllers.AddAgency).Methods("POST")
	r.HandleFunc("/agencies/edit/{id}", controllers.UpdateAgency).Methods("POST")
	r.HandleFunc("/agencies/{id}", controllers.GetOneAgency).Methods("GET")
	r.HandleFunc("/agencies/delete/{id}", controllers.DeleteAgency).Methods("GET")

	//Поставщики

	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
