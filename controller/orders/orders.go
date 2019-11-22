package orders

import (
	model "LSP/models"
	"LSP/util"
	"context"
	"encoding/json"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOrders(res http.ResponseWriter, req *http.Request) {
	var order model.Orders
	var orders []model.Orders
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("orders")
	cur, err := collection.Find(context.TODO(), bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()) {
		err := cur.Decode(&order)
		util.OnErr(err)
		orders = append(orders, order)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, orders)
}

func CreateOrders(res http.ResponseWriter, req *http.Request) {
	var order model.Orders
	err := json.NewDecoder(req.Body).Decode(&order)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("orders")
	collection.InsertOne(context.TODO(), order)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, order)
}

func GetOrders(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var order model.Orders
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("orders")
	util.OnErr(err)
	objid, err := primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	collection.FindOne(context.TODO(),bson.M{"_id":objid}).Decode(&order)
	util.ReturnRes(res, order)

	db.Disconnect(context.TODO())
}

// //GetAllOrders GetAllOrders
// func GetAllOrders(res http.ResponseWriter, req *http.Request) {
// 	var order model.Orders
// 	var orders []model.Orders
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from orders")
// 	for rows.Next() {
// 		rows.Scan(&order.Id, &order.Cart_id, &order.Qty, &order.Price, &order.Created_At)
// 		orders = append(orders, order)
// 	}
// 	util.ReturnRes(res, orders)
// }

// //CreateOrders CreateOrders
// func CreateOrders(res http.ResponseWriter, req *http.Request) {
// 	var order model.Orders
// 	err := json.NewDecoder(req.Body).Decode(&order)
// 	util.OnErr(err)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	query := "insert into orders values (''," + strconv.Itoa(order.Cart_id) + "," + strconv.Itoa(order.Qty) + "," + strconv.Itoa(order.Price) + ",'')"
// 	db.Exec(query)
// 	util.ReturnRes(res, order.Id)
// }

// func GetOrders(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	var order model.Orders
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from orders where id='" + id + "'")
// 	for rows.Next() {
// 		err := rows.Scan(&order.Id, &order.Cart_id, &order.Qty, &order.Price, &order.Created_At)
// 		util.OnErr(err)
// 	}
// 	util.ReturnRes(res, order)
// }
