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

func GetAllUnit(res http.ResponseWriter, req *http.Request) {
	log.Println("Requested")
	var unit model.Unit
	var unit_list []model.Unit
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("units")
	cur, err := collection.Find(context.TODO(), bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()) {
		err := cur.Decode(&unit)
		util.OnErr(err)
		unit_list = append(unit_list, unit)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, unit_list)
}

func DeleteUnit(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)
	objid,err:=primitive.ObjectIDFromHex(id)
	collection := db.Database("LSP").Collection("units")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, deleteResult.DeletedCount)
}

func CreateUnit(res http.ResponseWriter, req *http.Request) {
	log.Println("")
	var unit model.Unit
	err := json.NewDecoder(req.Body).Decode(&unit)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	// query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + string(password) + "','" + user.Role + "')"
	collection := db.Database("LSP").Collection("units")
	collection.InsertOne(context.TODO(), unit)

	db.Disconnect(context.TODO())
	util.ReturnRes(res, unit.Name)
}