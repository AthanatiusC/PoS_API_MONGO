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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllRole(res http.ResponseWriter, req *http.Request) {
	log.Println("Requested")
	var role model.Role
	var role_list []model.Role
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("roles")
	cur, err := collection.Find(context.TODO(), bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()) {
		err := cur.Decode(&role)
		util.OnErr(err)
		role_list = append(role_list, role)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, role_list)
}

func GetRole(res http.ResponseWriter, req *http.Request) {
	var role model.Role
	raw_param := mux.Vars(req)

	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)

	objid, err := primitive.ObjectIDFromHex(id)
	util.OnErr(err)

	collection := db.Database("LSP").Collection("roles")
	util.OnErr(err)
	err = collection.FindOne(context.TODO(), bson.M{"_id":objid}).Decode(&role)
	util.OnErr(err)
	db.Disconnect(context.TODO())
	util.ReturnRes(res, role)
}

func DeleteRole(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)
	objid,err:=primitive.ObjectIDFromHex(id)
	collection := db.Database("LSP").Collection("roles")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, deleteResult.DeletedCount)
}

func CreateRole(res http.ResponseWriter, req *http.Request) {
	log.Println("")
	var role model.Role
	err := json.NewDecoder(req.Body).Decode(&role)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	// query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + string(password) + "','" + user.Role + "')"
	collection := db.Database("LSP").Collection("roles")
	collection.InsertOne(context.TODO(), role)

	db.Disconnect(context.TODO())
	util.ReturnRes(res, role.Name)
}