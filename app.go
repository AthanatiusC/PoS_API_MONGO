package main

import (
	controller_carts "LSP/controller/cart"
	controller_customers "LSP/controller/customers"
	controller_orders "LSP/controller/orders"
	controller_products "LSP/controller/products"
	controller_profile "LSP/controller/profile"
	controller_users "LSP/controller/users"
	controller_role "LSP/controller/role"
	controller_unit "LSP/controller/unit"
	controller_category "LSP/controller/category"
	controller_currency "LSP/controller/currency"
	controller_minstock "LSP/controller/minstock"
	controller_transaction "LSP/controller/transaction"
	"github.com/rs/cors"
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
	// c := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:4200"},
	// })

	go router.HandleFunc("/", Home).Methods("GET")

	go router.HandleFunc("/users", controller_users.GetAllUser).Methods("GET")
	go router.HandleFunc("/users", controller_users.CreateUser).Methods("POST")
	go router.HandleFunc("/users/{id}", controller_users.GetUser).Methods("GET")
	go router.HandleFunc("/users/{username}", controller_users.UpdateUser).Methods("PUT")
	go router.HandleFunc("/users/delete/{id}", controller_users.DeleteUser).Methods("GET")
	go router.HandleFunc("/users/auth", controller_users.Auth).Methods("POST")

	go router.HandleFunc("/products", controller_products.CreateProduct).Methods("POST")
	go router.HandleFunc("/products", controller_products.GetAllProduct).Methods("GET")
	go router.HandleFunc("/products/{id}", controller_products.GetProduct).Methods("GET")
	go router.HandleFunc("/products/{id}", controller_products.UpdateProduct).Methods("POST")
	go router.HandleFunc("/products/delete/{id}", controller_products.DeleteProduct).Methods("GET")

	go router.HandleFunc("/orders", controller_orders.GetAllOrders).Methods("GET")
	go router.HandleFunc("/orders", controller_orders.CreateOrders).Methods("POST")
	go router.HandleFunc("/orders/{id}", controller_orders.GetOrders).Methods("GET")
	// // go router.HandleFunc("/orders/{id}", UpdateOrders).Methods("PUT")
	// // go router.HandleFunc("/orders/{id}", DeleteOrders).Methods("DELETE")

	go router.HandleFunc("/carts", controller_carts.CreateCarts).Methods("POST")
	go router.HandleFunc("/carts", controller_carts.GetAllCarts).Methods("GET")
	go router.HandleFunc("/carts/{id}", controller_carts.GetCarts).Methods("GET")
	go router.HandleFunc("/carts/{id}", controller_carts.UpdateCarts).Methods("PUT")
	go router.HandleFunc("/carts/delete/{id}", controller_carts.DeleteCarts).Methods("GET")
	go router.HandleFunc("/carts/today/", controller_carts.GetCartsToday).Methods("GET")
	go router.HandleFunc("/export", controller_carts.Export).Methods("GET")


	go router.HandleFunc("/customers", controller_customers.CreateCustomers).Methods("POST")
	go router.HandleFunc("/customers", controller_customers.GetAllCustomers).Methods("GET")
	go router.HandleFunc("/customers/{id}", controller_customers.GetCustomers).Methods("GET")
	go router.HandleFunc("/customers/{id}", controller_customers.UpdateCustomers).Methods("PUT")
	go router.HandleFunc("/customers/delete/{id}", controller_customers.DeleteCustomers).Methods("GET")

	go router.HandleFunc("/roles", controller_role.GetAllRole).Methods("GET")
	go router.HandleFunc("/roles", controller_role.CreateRole).Methods("POST")
	go router.HandleFunc("/roles/{id}", controller_role.GetRole).Methods("GET")
	go router.HandleFunc("/roles/delete/{id}", controller_role.DeleteRole).Methods("GET")

	go router.HandleFunc("/currency", controller_currency.GetAllCurrency).Methods("GET")
	go router.HandleFunc("/currency", controller_currency.CreateCurrency).Methods("POST")
	go router.HandleFunc("/currency/delete/{id}", controller_currency.DeleteCurrency).Methods("GET")

	go router.HandleFunc("/minstock", controller_minstock.GetAllCurrency).Methods("GET")
	go router.HandleFunc("/minstock", controller_minstock.CreateCurrency).Methods("POST")
	go router.HandleFunc("/minstock/delete/{id}", controller_minstock.DeleteCurrency).Methods("GET")

	go router.HandleFunc("/category", controller_category.GetAllCategory).Methods("GET")
	go router.HandleFunc("/category", controller_category.CreateCategory).Methods("POST")
	go router.HandleFunc("/category/delete/{id}", controller_category.DeleteCategory).Methods("GET")
	
	go router.HandleFunc("/unit", controller_unit.GetAllUnit).Methods("GET")
	go router.HandleFunc("/unit", controller_unit.CreateUnit).Methods("POST")
	go router.HandleFunc("/unit/delete/{id}", controller_unit.DeleteUnit).Methods("GET")

	go router.HandleFunc("/profile", controller_profile.GetProfile).Methods("GET")
	go router.HandleFunc("/profile", controller_profile.UpdateProfile).Methods("POST")
	go router.HandleFunc("/profile", controller_profile.UpdateProfile).Methods("PUT")

	go router.HandleFunc("/transaction", controller_transaction.CreateTransaction).Methods("POST")
	go router.HandleFunc("/transaction", controller_transaction.GetAllTransaction).Methods("GET")
	
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe("localhost:3030", handler))
}

//Home Display Category:Default
func Home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Go lang REST API</h1><br><span>Made with love by Lexi Anugrah</span>"))
}
