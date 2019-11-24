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

func GetAllCategory(res http.ResponseWriter, req *http.Request) {
	log.Println("Requested")
	var category model.Category
	var category_list []model.Category
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("category")
	cur, err := collection.Find(context.TODO(), bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()) {
		err := cur.Decode(&category)
		util.OnErr(err)
		category_list = append(category_list, category)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, category_list)
}

func DeleteCategory(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)
	objid,err:=primitive.ObjectIDFromHex(id)
	collection := db.Database("LSP").Collection("category")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, deleteResult.DeletedCount)
}

func CreateCategory(res http.ResponseWriter, req *http.Request) {
	log.Println("")
	var category model.Category
	err := json.NewDecoder(req.Body).Decode(&category)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	// query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + string(password) + "','" + user.Role + "')"
	collection := db.Database("LSP").Collection("category")
	collection.InsertOne(context.TODO(), category)

	db.Disconnect(context.TODO())
	util.ReturnRes(res, category.Name)
}