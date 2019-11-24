package products

import(
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"log"
	"net/http"
	//"strconv"
	"context"

	// "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	// "github.com/segmentio/ksuid"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTransaction(res http.ResponseWriter,req *http.Request){
	log.Println("")
	var transaction model.Transaction
	err:=json.NewDecoder(req.Body).Decode(&transaction)
	util.OnErr(err)

	db,err:=util.ConnectMongo()
	util.OnErr(err)
	_,err=db.Database("LSP").Collection("carts").DeleteMany(context.TODO(),bson.M{})
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("transaction")
	collection.InsertOne(context.TODO(),transaction)
	db.Disconnect(context.TODO())

	util.ReturnRes(res,transaction.User_id)
}

func GetAllTransaction(res http.ResponseWriter,req *http.Request){
	var transaction model.Transaction
	var transactions []model.Transaction
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("transaction")
	cur,err:=collection.Find(context.TODO(),bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()){
		err:=cur.Decode(&transaction)
		util.OnErr(err)
		transactions=append(transactions,transaction)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res,transactions)
}