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
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProfile(res http.ResponseWriter,req *http.Request){
	log.Println("")
	var profile model.Profile
	err:=json.NewDecoder(req.Body).Decode(&profile)
	util.OnErr(err)
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("profile")
	collection.InsertOne(context.TODO(),profile)
	log.Println(profile)
	db.Disconnect(context.TODO())
	if profile.Description != ""{
		util.ReturnRes(res,profile)
		return
	}
	util.ReturnRes(res,nil)
}

func GetProfile(res http.ResponseWriter,req *http.Request){
	//varproducts[]model.Products
	var profile model.Profile
	// raw_param:=mux.Vars(req)
	// id:=raw_param["id"]
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	// objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("profile")
	util.OnErr(err)

	collection.FindOne(context.TODO(),bson.M{}).Decode(&profile)

	db.Disconnect(context.TODO())

	if profile.Description != ""{
		util.ReturnRes(res,profile)
		return
	}
	util.ReturnRes(res,nil)
}

func UpdateProfile(res http.ResponseWriter,req *http.Request){
	// raw_param:=mux.Vars(req)
	// id:=raw_param["id"]
	var profile model.Profile
	err:=json.NewDecoder(req.Body).Decode(&profile)
	log.Pritnln(profile)
	util.OnErr(err)
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("profile")
	data:=bson.D{{Key:"$set",Value:profile}}
	// objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	result,err:=collection.UpdateOne(context.TODO(),bson.M{},data)
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res,result.ModifiedCount)
}