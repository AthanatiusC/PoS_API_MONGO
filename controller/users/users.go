package users

import (
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"context"
	"log"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// var db *mongo.Client

func CreateUser(res http.ResponseWriter, req *http.Request) {
	log.Println("")
	var user model.Users
	err := json.NewDecoder(req.Body).Decode(&user)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	util.OnErr(err)

	user.Password = string(password)

	if len(user.Username) == 0{
		util.ReturnRes(res, nil)
		return
	}

	// query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + string(password) + "','" + user.Role + "')"
	collection := db.Database("LSP").Collection("users")
	collection.InsertOne(context.TODO(), user)
	db.Disconnect(context.TODO())
	util.ReturnRes(res, user.Name)
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	var users model.Users
	raw_param := mux.Vars(req)

	id := raw_param["id"]
	db, err := util.ConnectMongo()
	util.OnErr(err)

	objid, err := primitive.ObjectIDFromHex(id)
	util.OnErr(err)

	collection := db.Database("LSP").Collection("users")
	util.OnErr(err)
	err = collection.FindOne(context.TODO(), bson.M{"_id":objid}).Decode(&users)
	util.OnErr(err)

	db.Disconnect(context.TODO())
	util.ReturnRes(res, users)
}

func GetAllUser(res http.ResponseWriter, req *http.Request) {
	log.Println("Requested")
	var users model.Users
	var user_list []model.Users
	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()) {
		err := cur.Decode(&users)
		util.OnErr(err)
		user_list = append(user_list, users)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res, user_list)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	raw_param := mux.Vars(req)
	id := raw_param["id"]
	objid,err:=primitive.ObjectIDFromHex(id)

	db, err := util.ConnectMongo()
	util.OnErr(err)

	collection := db.Database("LSP").Collection("users")
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, deleteResult.DeletedCount)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var user model.Users
	raw_param := mux.Vars(req)
	uname := raw_param["username"]
	err := json.NewDecoder(req.Body).Decode(&user)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	if len(user.Username) == 0{
		util.ReturnRes(res, nil)
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	collection := db.Database("LSP").Collection("users")
	data := bson.D{ {Key: "$set", Value: user} }
	result,err := collection.UpdateOne(context.TODO(),bson.M{"username":uname},data)
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res, result.ModifiedCount)
}

func Auth(res http.ResponseWriter, req *http.Request) {
	var user model.Users
	var userauth model.Users
	err := json.NewDecoder(req.Body).Decode(&user)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	collection := db.Database("LSP").Collection("users")
	err = collection.FindOne(context.TODO(), bson.M{"username":user.Username}).Decode(&userauth)
	util.OnErr(err)
	
	// password:= hashAndSalt()
	ismatch := comparePasswords(userauth.Password,[]byte(user.Password))

	db.Disconnect(context.TODO())
	if ismatch == true{
		util.ReturnRes(res,userauth)
	}else{
		util.ReturnRes(res,nil)
	}
}


func hashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {    // Since we'll be getting the hashed password from the DB it
	byteHash := []byte(hashedPwd)    
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        return false
    }
	return true
}


// func GetUser(res http.ResponseWriter, req *http.Request) {
// 	var users model.Users
// 	raw_param := mux.Vars(req)
// 	param := raw_param["id"]
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from users where id = '" + param + "'")
// 	util.OnErr(err)
// 	for rows.Next() {

// 		err := rows.Scan(&users.Id, &users.Name, &users.Username, &users.Password, &users.Role)
// 		util.OnErr(err)
// 	}
// 	util.ReturnRes(res, users)
// }

// func GetAllUser(res http.ResponseWriter, req *http.Request) {
// 	var users model.Users
// 	var user_list []model.Users
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	rows, err := db.Query("select * from users")
// 	util.OnErr(err)
// 	for rows.Next() {
// 		err := rows.Scan(&users.Id, &users.Name, &users.Username, &users.Password, &users.Role)
// 		util.OnErr(err)
// 		user_list = append(user_list, users)
// 	}
// 	util.ReturnRes(res, user_list)
// }

// func CreateUser(res http.ResponseWriter, req *http.Request) {
// 	var user model.Users
// 	err := json.NewDecoder(req.Body).Decode(&user)
// 	util.OnErr(err)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	util.OnErr(err)
// 	query := "insert into users values ('','" + user.Name + "','" + user.Username + "','" + string(password) + "','" + user.Role + "')"
// 	db.Exec(query)
// 	util.ReturnRes(res, user.Name)
// }

// func UpdateUser(res http.ResponseWriter, req *http.Request) {
// 	var user model.Users
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	err := json.NewDecoder(req.Body).Decode(&user)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	query := "update users set name='" + user.Name + "',username='" + user.Username + "',password='" + string(password) + "',role='" + user.Role + "' where id='" + id + "'"
// 	log.Println(query)
// 	_, err = db.Query(query)
// 	util.OnErr(err)
// 	util.ReturnRes(res, user.Name)
// }

// func Auth(res http.ResponseWriter, req *http.Request) {
// 	var user model.Users
// 	var userauth model.Users
// 	err := json.NewDecoder(req.Body).Decode(&user)
// 	util.OnErr(err)
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	query := "select * from users where username='" + user.Username + "'"
// 	rows, err := db.Query(query)
// 	util.OnErr(err)
// 	for rows.Next() {
// 		err := rows.Scan(&userauth.Id, &userauth.Name, &userauth.Username, &userauth.Password, &userauth.Role)
// 		util.OnErr(err)
// 	}

// 	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	err = bcrypt.CompareHashAndPassword([]byte(userauth.Password), password)
// 	if err != nil {
// 		util.ReturnRes(res, userauth)
// 	} else {
// 		util.ReturnRes(res, nil)
// 	}
// }

// func DeleteUser(res http.ResponseWriter, req *http.Request) {
// 	raw_param := mux.Vars(req)
// 	id := raw_param["id"]
// 	db, err := util.ConnectDB()
// 	util.OnErr(err)
// 	_, err = db.Query("delete from users where id='" + id + "'")
// 	util.OnErr(err)
// 	util.ReturnRes(res, nil)
// }
