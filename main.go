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

	//Бухгалтерия
	r.HandleFunc("/buh", controllers.GetAllBuh).Methods("GET")
	r.HandleFunc("/buh/add", controllers.AddBuh).Methods("POST")
	r.HandleFunc("/buh/edit/{id}", controllers.UpdateBuh).Methods("POST")
	r.HandleFunc("/buh/{id}", controllers.GetOneBuh).Methods("GET")
	r.HandleFunc("/buh/delete/{id}", controllers.DeleteBuh).Methods("GET")

	//Сотрудники
	r.HandleFunc("/employees", controllers.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/add", controllers.AddEmployee).Methods("POST")
	r.HandleFunc("/employees/edit/{id}", controllers.UpdateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", controllers.GetEmployee).Methods("GET")
	r.HandleFunc("/employees/delete/{id}", controllers.DeleteEmployee).Methods("GET")

	//Закупка
	r.HandleFunc("/purchase", controllers.GetAllPurchases).Methods("GET")
	r.HandleFunc("/purchase/add", controllers.AddPurchase).Methods("POST")
	r.HandleFunc("/purchase/edit/{id}", controllers.UpdatePurchase).Methods("POST")
	r.HandleFunc("/purchase/{id}", controllers.GetOnePurchase).Methods("GET")
	r.HandleFunc("/purchase/delete/{id}", controllers.DeletePurchase).Methods("GET")

	//Продажа
	r.HandleFunc("/sales", controllers.GetAllSales).Methods("GET")
	r.HandleFunc("/sales/add", controllers.AddSale).Methods("POST")
	r.HandleFunc("/sales/edit/{id}", controllers.UpdateSale).Methods("POST")
	r.HandleFunc("/sales/{id}", controllers.GetOneSale).Methods("GET")
	r.HandleFunc("/sales/delete/{id}", controllers.DeleteSale).Methods("GET")

	//Чек
	r.HandleFunc("/cheque", controllers.GetAllCheques).Methods("GET")
	r.HandleFunc("/cheque/add", controllers.AddCheque).Methods("POST")
	r.HandleFunc("/cheque/{id}", controllers.GetOneCheque).Methods("GET")
	r.HandleFunc("/cheque/delete/{id}", controllers.DeleteCheque).Methods("GET")

	//Поставка
	r.HandleFunc("/supplies", controllers.GetAllSupplies).Methods("GET")
	r.HandleFunc("/supplies/add", controllers.AddSupply).Methods("POST")
	r.HandleFunc("/supplies/{id}", controllers.GetOneSupply).Methods("GET")
	r.HandleFunc("/supplies/delete/{id}", controllers.DeleteSupply).Methods("GET")

	//Авторизация
	r.HandleFunc("/login", controllers.SignIn).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
