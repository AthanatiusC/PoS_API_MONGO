package carts

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func GetAllCarts(res http.ResponseWriter, req *http.Request) {
	var cart model.Carts
	var carts []model.Carts
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	collection:=db.Database("LSP").Collection("carts")
	cur,err:=collection.Find(context.TODO(),bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()){
		err:=cur.Decode(&cart)
		util.OnErr(err)
		carts=append(carts,cart)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, carts)
}

func CreateCarts(res http.ResponseWriter, req *http.Request) {
	var cart model.Carts
	log.Println("")
	err := json.NewDecoder(req.Body).Decode(&cart)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	cart.Invoice = string(strconv.FormatFloat(rand.Float64(),'E', -1, 64) + time.Now().Format("/01/Jan/2006/150405"))
	cart.Created_At = time.Now().Format("2006-01-02 3:4:5 PM")
	cart.Updated_At = time.Now().Format("2006-01-02 3:4:5 PM")
	collection:=db.Database("LSP").Collection("carts")
	collection.InsertOne(context.TODO(),cart)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, cart)
}

func GetCarts(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var cart model.Carts
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	collection:=db.Database("LSP").Collection("carts")
	util.OnErr(err)
	collection.FindOne(context.TODO(),bson.M{"_id":objid}).Decode(&cart)

	db.Disconnect(context.TODO())

	util.ReturnRes(res, cart)
}

func UpdateCarts(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	var cart model.Carts
	err := json.NewDecoder(req.Body).Decode(&cart)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)

	cart.Updated_At = time.Now().Format("2006-01-02 3:4:5 PM")
	log.Println(cart)

	collection:=db.Database("LSP").Collection("carts")
	data:=bson.D{{Key:"$set",Value:cart}}
	objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	result,err:=collection.UpdateOne(context.TODO(),bson.M{"_id":objid},data)
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, result.ModifiedCount)
}

func DeleteCarts(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)

	objid,err:=primitive.ObjectIDFromHex(id)
	collection:=db.Database("LSP").Collection("carts")
	deleteResult,err:=collection.DeleteOne(context.TODO(),bson.M{"_id":objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, deleteResult.DeletedCount)
}

// func GetCarts(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	var cart model.Carts
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from carts where id='" + id + "'")
// 	util.OnErr(err)
// 	for rows.Next() {
// 		err := rows.Scan(&cart.Id, &cart.Invoice, &cart.Customer_id, &cart.Total, &cart.Created_At, &cart.Updated_At)
// 		util.OnErr(err)
// 	}
// 	util.ReturnRes(res, cart)
// }

// //CreateCarts CreateCarts
// func CreateCarts(res http.ResponseWriter, req *http.Request) {
// 	var cart model.Carts
// 	err := json.NewDecoder(req.Body).Decode(&cart)
// 	util.OnErr(err)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	invoice := string(strconv.Itoa(rand.Intn(1000)) + time.Now().Format("/01/Jan/2006/150405"))
// 	query := "insert into carts values ('','" + invoice + "','" + strconv.Itoa(cart.Customer_id) + "','" + strconv.Itoa(cart.Total) + "','','')"
// 	db.Exec(query)
// 	util.ReturnRes(res, cart.Id)
// }

// //GetAllCarts Category:Carts
// func GetAllCarts(res http.ResponseWriter, req *http.Request) {
// 	var cart model.Carts
// 	var carts []model.Carts
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from carts")
// 	for rows.Next() {
// 		err := rows.Scan(&cart.Id, &cart.Invoice, &cart.Customer_id, &cart.Total, &cart.Created_At, &cart.Updated_At)
// 		util.OnErr(err)
// 		carts = append(carts, cart)
// 	}
// 	util.ReturnRes(res, carts)
// }

// //GetCarts GetCarts
// func GetCarts(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	var cart model.Carts
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from carts where id='" + id + "'")
// 	util.OnErr(err)
// 	for rows.Next() {
// 		err := rows.Scan(&cart.Id, &cart.Invoice, &cart.Customer_id, &cart.Total, &cart.Created_At, &cart.Updated_At)
// 		util.OnErr(err)
// 	}
// 	util.ReturnRes(res, cart)
// }

// //UpdateCarts UpdateCarts
// func UpdateCarts(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	var cart model.Carts
// 	err := json.NewDecoder(req.Body).Decode(&cart)
// 	util.OnErr(err)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	query := "update carts set total='" + strconv.Itoa(cart.Total) + "',updated_at='' where id='" + id + "'"
// 	log.Println(query)
// 	_, err = db.Query(query)
// 	util.OnErr(err)
// 	util.ReturnRes(res, nil)
// }

// //DeleteCarts DeleteCarts
// func DeleteCarts(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	_, err = db.Query("delete from carts where id='" + id + "'")
// 	util.OnErr(err)
// 	util.ReturnRes(res, nil)
// }
