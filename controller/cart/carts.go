package carts

import (
	model "LSP/models"
	"bufio"
	"LSP/util"
	"encoding/json"
	"log"
	// "math/rand"
	"net/http"
	"strconv"
	"time"
	"context"
	"os"
	"io"
	"fmt"
	// "encoding/csv"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func Export(writer http.ResponseWriter, req *http.Request) {
	CSV()
	Filename := "Data.csv"
	if Filename == "" {
		//Get not set, send a 400 bad request
		http.Error(writer, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + Filename)

	//Check if file exists and open
	Openfile, err := os.Open("D:\\Server\\Lexi\\Lexi.csv")
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(writer, "File not found.", 404)
		return
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
	return
}

func CSV(){
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

	var conf []string

	for _,data :=range carts{
		// qty:= 
		conf = append(conf,data.Invoice+","+strconv.Itoa(data.Qty)+","+strconv.Itoa(data.Total)+","+data.Created_At.String()+","+data.Updated_At.String())
	}
	log.Println("Exporting..")
	f, err := os.Create("D:\\Server\\Lexi\\Lexi.csv")
	defer f.Close()
	util.OnErr(err)
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	for _, line := range conf {
		_, err = writer.WriteString(line + "\n")
		util.OnErr(err)
	}
	log.Println("Completed!")
}

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
	var product model.Products
	log.Println("")
	err := json.NewDecoder(req.Body).Decode(&cart)
	util.OnErr(err)
	db, err := util.ConnectMongo()
	util.OnErr(err)
	id := ksuid.New()
	cart.Invoice = string(id.String() + time.Now().Format("/01/Jan/2006/150405"))
	cart.Created_At = time.Now()
	cart.Updated_At = time.Now()
	cart.Ppn = 1450
	cart.Percentage = 2
	cart.Price = product.Price * cart.Qty + cart.Ppn
	// objid,err:=primitive.ObjectIDFromHex(cart.Product_id)
	util.OnErr(err)
	col2:=db.Database("LSP").Collection("products")
	util.OnErr(err)
	col2.FindOne(context.TODO(),bson.M{"_id":cart.Product_id}).Decode(&product)

	if cart.Disc != 0{
		cart.Total = ((product.Price * cart.Qty + cart.Ppn)+ (product.Price * cart.Qty + cart.Ppn)*cart.Percentage/100)*cart.Disc/100
	}else{
		cart.Total = (product.Price * cart.Qty + cart.Ppn)+ (product.Price * cart.Qty + cart.Ppn)*cart.Percentage/100
	}

	log.Println(product.Price)
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

func GetCartsToday(res http.ResponseWriter, req *http.Request) {
	var cart model.Carts
	db, err := util.ConnectMongo()
	util.OnErr(err)
	
	// objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	collection:=db.Database("LSP").Collection("carts")
	util.OnErr(err)
	collection.FindOne(context.TODO(),
	bson.M{"created_at":bson.M{
		"$gt":time.Now().Add(-24*time.Hour),"$lt":time.Now().Format("2006-01-02 3:4:5 PM")}},).Decode(&cart)

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

	cart.Updated_At = time.Now()
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
