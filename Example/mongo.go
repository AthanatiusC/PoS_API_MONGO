package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

type Schedules struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Teacher  string             `json:"teacher,omitempty" bson:"teacher,omitempty"`
	Day      string             `json:"day,omitempty" bson:"day,omitempty"`
	Duration int                `json:"duration,omitempty" bson:"duration,omitempty"`
}

func main() {
	log.Println("Booting...")
	context, _ := context.WithTimeout(context.Background(), time.Second*10)

	clientopt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context, clientopt)
	onError(err)

	router := mux.NewRouter()

	router.HandleFunc("/schedule/{day}", getScheduleToday).Methods("GET")
	router.HandleFunc("/schedule", postSchedule).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getScheduleToday(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	param := mux.Vars(req)
	day, _ := param["day"]
	var Schedule Schedules
	collection := client.Database("School").Collection("Schedule")
	context, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := collection.FindOne(context, Schedules{Day: day}).Decode(&Schedule)
	onError(err)
	json.NewEncoder(res).Encode(Schedule)
}

func postSchedule(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var schedule Schedules
	_ = json.NewDecoder(req.Body).Decode(&schedule)
	collection := client.Database("School").Collection("Schedule")
	context, _ := context.WithTimeout(context.Background(), time.Second*10)
	status, err := collection.InsertOne(context, schedule)
	onError(err)
	json.NewEncoder(res).Encode(status)
}

func onError(err error) {
	if err != nil {
		log.Println(err)
	}
}
