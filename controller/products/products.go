package products

import(
	model "LSP/models"
	"LSP/util"
	"encoding/json"
	"log"
	"net/http"
	//"strconv"
	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(res http.ResponseWriter,req *http.Request){
	log.Println("")
	var product model.Products
	err:=json.NewDecoder(req.Body).Decode(&product)
	util.OnErr(err)
	if len(product.Curr) == 0{
		util.ReturnRes(res,nil)
		return
	}
db,err:=util.ConnectMongo()
	util.OnErr(err)
	id := ksuid.New()
	product.Barcode = id.String()

	collection:=db.Database("LSP").Collection("products")
	collection.InsertOne(context.TODO(),product)
	db.Disconnect(context.TODO())

	util.ReturnRes(res,product.Name)
}

func GetProduct(res http.ResponseWriter,req *http.Request){
	//varproducts[]model.Products
	var product model.Products
	raw_param:=mux.Vars(req)
	id:=raw_param["id"]
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("products")
	util.OnErr(err)

	collection.FindOne(context.TODO(),bson.M{"_id":objid}).Decode(&product)
	
	db.Disconnect(context.TODO())

	util.ReturnRes(res,product)
}

func GetAllProduct(res http.ResponseWriter,req *http.Request){
	var product model.Products
	var products[] model.Products
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	collection:=db.Database("LSP").Collection("products")
	cur,err:=collection.Find(context.TODO(),bson.M{})
	util.OnErr(err)
	for cur.Next(context.TODO()){
		err:=cur.Decode(&product)
		util.OnErr(err)
		products=append(products,product)
	}
	db.Disconnect(context.TODO())

	util.ReturnRes(res,products)
}

func UpdateProduct(res http.ResponseWriter,req *http.Request){
	raw_param:=mux.Vars(req)
	id:=raw_param["id"]
	var product model.Products
	err:=json.NewDecoder(req.Body).Decode(&product)
	util.OnErr(err)
	db,err:=util.ConnectMongo()
	util.OnErr(err)

	if len(product.Curr) == 0{
		util.ReturnRes(res,nil)
		return
	}

	collection:=db.Database("LSP").Collection("products")
	data:=bson.D{{Key:"$set",Value:product}}
	objid,err:=primitive.ObjectIDFromHex(id)
	util.OnErr(err)
	result,err:=collection.UpdateOne(context.TODO(),bson.M{"_id":objid},data)
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res,result.ModifiedCount)
}

func DeleteProduct(res http.ResponseWriter,req *http.Request){
	raw_param:=mux.Vars(req)
	id:=raw_param["id"]
	db,err:=util.ConnectMongo()
	util.OnErr(err)
	objid,err:=primitive.ObjectIDFromHex(id)

	collection:=db.Database("LSP").Collection("products")
	deleteResult,err:=collection.DeleteOne(context.TODO(),bson.M{"_id":objid})
	util.OnErr(err)
	db.Disconnect(context.TODO())

	util.ReturnRes(res,deleteResult.DeletedCount)
}

////CreateProductCategory:Products
//funcCreateProduct(reshttp.ResponseWriter,req*http.Request){
//	varproductmodel.Products
//	err:=json.NewDecoder(req.Body).Decode(&product)
//	util.OnErr(err)
//	db,err:=util.ConnectDB()
//	util.OnErr(err)
//	query:="insertintoproductvalues('','"+product.Name+"',"+strconv.Itoa(product.Price)+","+strconv.Itoa(product.Stock)+")"
//	db.Exec(query)
//	util.ReturnRes(res,product.Name)
//}

////GetProductGetProduct
//funcGetProduct(reshttp.ResponseWriter,req*http.Request){
//	varproductmodel.Products
//	raw_param:=mux.Vars(req)
//	id:=raw_param["id"]
//	db,err:=util.ConnectDB()
//	util.OnErr(err)
//	rows,err:=db.Query("select*fromproductwhereid='"+id+"'")
//	util.OnErr(err)
//	forrows.Next(){
//		err:=rows.Scan(rows.Scan(&product.Id,&product.Name,&product.Price,&product.Stock))
//		util.OnErr(err)
//	}
//	util.ReturnRes(res,product)
//}

////GetAllProductGetAllProduct
//funcGetAllProduct(reshttp.ResponseWriter,req*http.Request){
//	varproductmodel.Products
//	varproducts[]model.Products
//	db,err:=util.ConnectDB()
//	util.OnErr(err)
//	rows,err:=db.Query("select*fromproduct")
//	forrows.Next(){
//		err:=rows.Scan(&product.Id,&product.Name,&product.Price,&product.Stock)
//		util.OnErr(err)
//		products=append(products,product)
//	}
//	util.ReturnRes(res,products)
//}

////UpdateProductUpdateProduct
//funcUpdateProduct(reshttp.ResponseWriter,req*http.Request){
//	raw_param:=mux.Vars(req)
//	id:=raw_param["id"]
//	varproductmodel.Products
//	err:=json.NewDecoder(req.Body).Decode(&product)
//	util.OnErr(err)
//	db,err:=util.ConnectDB()
//	util.OnErr(err)
//	query:="updateproductsetname='"+product.Name+"',price="+strconv.Itoa(product.Price)+",stock="+strconv.Itoa(product.Stock)+"whereid='"+id+"'"
//	log.Println(query)
//	_,err=db.Query(query)
//	util.OnErr(err)
//	util.ReturnRes(res,nil)
//}

////DeleteProductDeleteProduct
//funcDeleteProduct(reshttp.ResponseWriter,req*http.Request){
//	raw_param:=mux.Vars(req)
//	id:=raw_param["id"]
//	db,err:=util.ConnectDB()
//	util.OnErr(err)
//	_,err=db.Query("deletefromproductwhereid='"+id+"'")
//	util.OnErr(err)
//	util.ReturnRes(res,nil)
//}
