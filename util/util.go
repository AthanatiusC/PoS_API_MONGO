package util

import (
	// "database/sql"
	"context"
	"time"
	"encoding/json"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func ConnectSQL() (database *sql.DB, err error) {
// 	db, err := sql.Open("mysql", "root@tcp(localhost)/lsp")
// 	return db, err
// }

func ConnectMongo() (client *mongo.Client,err error) {
	context, _ := context.WithTimeout(context.Background(), time.Second*10)
	clientopt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context, clientopt)
	return client,err
}

func ReturnRes(res http.ResponseWriter, value interface{}) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(value)
}

func OnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
