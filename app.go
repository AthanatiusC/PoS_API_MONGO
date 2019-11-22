package main

import (
	controller_carts "LSP/controller/cart"
	controller_customers "LSP/controller/customers"
	controller_orders "LSP/controller/orders"
	controller_products "LSP/controller/products"
	controller_users "LSP/controller/users"
	"LSP/util"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	func() {
		db, err := util.ConnectMongo()
		util.OnErr(err)
		log.Println("Connection Established")
		context, _ := context.WithTimeout(context.Background(), time.Second*10)
		db.Disconnect(context)
	}()

	router := mux.NewRouter()
	go router.HandleFunc("/", Home).Methods("GET")

	go router.HandleFunc("/users", controller_users.GetAllUser).Methods("GET")
	go router.HandleFunc("/users", controller_users.CreateUser).Methods("POST")
	go router.HandleFunc("/users/{id}", controller_users.GetUser).Methods("GET")
	go router.HandleFunc("/users/{username}", controller_users.UpdateUser).Methods("PUT")
	go router.HandleFunc("/users/{username}", controller_users.DeleteUser).Methods("DELETE")
	go router.HandleFunc("/users/auth", controller_users.Auth).Methods("POST")

	go router.HandleFunc("/products", controller_products.CreateProduct).Methods("POST")
	go router.HandleFunc("/products", controller_products.GetAllProduct).Methods("GET")
	go router.HandleFunc("/products/{id}", controller_products.GetProduct).Methods("GET")
	go router.HandleFunc("/products/{id}", controller_products.UpdateProduct).Methods("PUT")
	go router.HandleFunc("/products/{id}", controller_products.DeleteProduct).Methods("DELETE")

	go router.HandleFunc("/orders", controller_orders.GetAllOrders).Methods("GET")
	go router.HandleFunc("/orders", controller_orders.CreateOrders).Methods("POST")
	go router.HandleFunc("/orders/{id}", controller_orders.GetOrders).Methods("GET")
	// // go router.HandleFunc("/orders/{id}", UpdateOrders).Methods("PUT")
	// // go router.HandleFunc("/orders/{id}", DeleteOrders).Methods("DELETE")

	go router.HandleFunc("/carts", controller_carts.CreateCarts).Methods("POST")
	go router.HandleFunc("/carts", controller_carts.GetAllCarts).Methods("GET")
	go router.HandleFunc("/carts/{id}", controller_carts.GetCarts).Methods("GET")
	go router.HandleFunc("/carts/{id}", controller_carts.UpdateCarts).Methods("PUT")
	go router.HandleFunc("/carts/{id}", controller_carts.DeleteCarts).Methods("DELETE")

	router.HandleFunc("/customers", controller_customers.CreateCustomers).Methods("POST")
	router.HandleFunc("/customers", controller_customers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controller_customers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controller_customers.UpdateCustomers).Methods("PUT")
	router.HandleFunc("/customers/{id}", controller_customers.DeleteCustomers).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:3030", router))
}

//Home Display Category:Default
func Home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Go lang REST API</h1><br><span>Made with love by Lexi Anugrah</span>"))
}
