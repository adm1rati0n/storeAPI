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
	//r.Handle("/api", controllers.ValidateJWT(controllers.GetPosts))
	//r.HandleFunc("/jwt", controllers.GetJWT)

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
	r.HandleFunc("/suppliers", controllers.GetAllSuppliers).Methods("GET")
	r.HandleFunc("/suppliers/add", controllers.AddSupplier).Methods("POST")
	r.HandleFunc("/suppliers/edit/{id}", controllers.UpdateSupplier).Methods("POST")
	r.HandleFunc("/suppliers/{id}", controllers.GetOneSupplier).Methods("GET")
	r.HandleFunc("/suppliers/delete/{id}", controllers.DeleteSupplier).Methods("GET")

	//Должности
	r.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
