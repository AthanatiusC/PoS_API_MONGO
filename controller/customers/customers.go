package customers

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"net/http"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
)

func CreateCustomers(res http.ResponseWriter, req *http.Request) {
	log.Println("")
	var customer model.Customers
	err := json.NewDecoder(req.Body).Decode(&customer)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	collection:=db.Database("LSP").Collection("customers")
	collection.InsertOne(context.TODO(),customer)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, customer.Name)
}

func GetAllCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	var customer_list []model.Customers
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("customers")
	cur,err:=collection.Find(context.TODO(),bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()){
		err:=cur.Decode(&customer)
		util.OnErr(err)
		customer_list=append(customer_list,customer)
	}
	
	db.Disconnect(context.TODO())
	util.ReturnRes(res, customer_list)
}

func GetCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	raw_param := mux.Vars(req)
	param := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	objid,err:=primitive.ObjectIDFromHex(param)
	util.OnErr(err)
	collection:=db.Database("LSP").Collection("customers")
	util.OnErr(err)
	collection.FindOne(context.TODO(),bson.M{"_id":objid}).Decode(&customer)

	util.ReturnRes(res, customer)
}

func UpdateCustomers(res http.ResponseWriter, req *http.Request) {
	var customer model.Customers
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	err := json.NewDecoder(req.Body).Decode(&customer)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	collection:=db.Database("LSP").Collection("customers")
	data:=bson.D{{Key:"$set",Value:customer}}
	objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	result,err:=collection.UpdateOne(context.TODO(),bson.M{"_id":objid},data)
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, result.ModifiedCount)
}

func DeleteCustomers(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	objid,err:=primitive.ObjectIDFromHex(id)
	collection:=db.Database("LSP").Collection("customers")
	deleteResult,err:=collection.DeleteOne(context.TODO(),bson.M{"_id":objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, deleteResult.DeletedCount)
}


// func CreateCustomers(res http.ResponseWriter, req *http.Request) {
// 	var customer model.Customers
// 	err := json.NewDecoder(req.Body).Decode(&customer)
// 	util.OnErr(err)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	query := "insert into customers values ('','" + customer.Email + "','" + customer.Name + "','" + customer.Adress + "','" + customer.Phone + "')"
// 	db.Exec(query)
// 	util.ReturnRes(res, customer.Name)
// }

// func GetAllCustomers(res http.ResponseWriter, req *http.Request) {
// 	var customer model.Customers
// 	var customer_list []model.Customers
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from customers")
// 	util.OnErr(err)
// 	for rows.Next() {
// 		err := rows.Scan(&customer.Id, &customer.Email, &customer.Name, &customer.Adress, &customer.Phone)
// 		util.OnErr(err)
// 		customer_list = append(customer_list, customer)
// 	}
// 	util.ReturnRes(res, customer_list)
// }

// func GetCustomers(res http.ResponseWriter, req *http.Request) {
// 	var customer model.Customers
// 	raw_param := mux.Vars(req)
// 	param := raw_param["id"]
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from customers where id = '" + param + "'")
// 	util.OnErr(err)
// 	for rows.Next() {
// 		err := rows.Scan(&customer.Id, &customer.Email, &customer.Name, &customer.Adress, &customer.Phone)
// 		util.OnErr(err)
// 	}
// 	util.ReturnRes(res, customer)
// }

// func UpdateCustomers(res http.ResponseWriter, req *http.Request) {
// 	var customer model.Customers
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	err := json.NewDecoder(req.Body).Decode(&customer)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	query := "update customers set email='" + customer.Email + "',name='" + customer.Name + "',adress='" + customer.Adress + "',phone='" + customer.Phone + "' where id='" + id + "'"
// 	_, err = db.Query(query)
// 	util.OnErr(err)
// 	util.ReturnRes(res, customer.Name)
// }

// func DeleteCustomers(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	_, err = db.Query("delete from customers where id='" + id + "'")
// 	util.OnErr(err)
// 	util.ReturnRes(res, nil)
// }
