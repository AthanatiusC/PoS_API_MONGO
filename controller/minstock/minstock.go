package products

import(
	model "LSP/models"
	"LSP/util"
	"context"
	"log"
	"net/http"
	"encoding/json"
	// "time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllCurrency(res http.ResponseWriter, req *http.Request) {
	log.Println("Requested")
	var currency model.MinStock
	var currency_list []model.MinStock
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("minstock")
	cur, err := collection.Find(context.TODO(), bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()) {
		err := cur.Decode(&currency)
		util.OnErr(err)
		currency_list = append(currency_list, currency)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, currency_list)
}

func DeleteCurrency(res http.ResponseWriter, req *http.Request) {
	raw_param:=mux.Vars(req)
	id:=raw_param["id"]
	db,err:=util.ConnectMongo()
	util.OnErr(err)
	objid,err:=primitive.ObjectIDFromHex(id)

	collection:=db.Database("LSP").Collection("minstock")
	deleteResult,err:=collection.DeleteOne(context.TODO(),bson.M{"_id":objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res,deleteResult.DeletedCount)
}

func CreateCurrency(res http.ResponseWriter, req *http.Request) {
	log.Println("")
	var currency model.MinStock
	err := json.NewDecoder(req.Body).Decode(&currency)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	// query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + string(password) + "','" + user.Role + "')"
	collection := db.Database("LSP").Collection("minstock")
	collection.InsertOne(context.TODO(), currency)

	db.Disconnect(context.TODO())
	util.ReturnRes(res, currency)
}