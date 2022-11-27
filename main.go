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
	r.HandleFunc("/posts/add", controllers.AddPost).Methods("POST")
	r.HandleFunc("/posts/edit/{id}", controllers.UpdatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", controllers.GetOnePost).Methods("GET")
	r.HandleFunc("/posts/delete/{id}", controllers.DeletePost).Methods("GET")

	//Пользователи
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/add", controllers.AddUser).Methods("POST")
	r.HandleFunc("/users/edit/{id}", controllers.UpdateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetOneUser).Methods("GET")
	r.HandleFunc("/users/delete/{id}", controllers.DeleteUser).Methods("GET")

	//Реклама
	r.HandleFunc("/ads", controllers.GetAllAds).Methods("GET")
	r.HandleFunc("/ads/add", controllers.AddAd).Methods("POST")
	r.HandleFunc("/ads/edit/{id}", controllers.UpdateAd).Methods("POST")
	r.HandleFunc("/ads/{id}", controllers.GetOneAd).Methods("GET")
	r.HandleFunc("/ads/delete/{id}", controllers.DeleteAd).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
