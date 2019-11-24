package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


type Users struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
	Contact string `json:"contact"`
}

type Transaction struct{
	Id    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	History []Carts `json:history`
	Money int `json:money`
	Exchange int `json:exchange`
	User_id primitive.ObjectID `json:"user_id"`
}

type Products struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Price int                `json:"price"`
	Stock int                `json:"stock"`
	Barcode string `json:"barcode"`
	Classification string`json:"classification"`
	Curr string`json:"curr"`
	SellPrice int`json:"sellprice"`
	BuyPrice int`json:"buyprice"`
	Unit string`json:"unit"`
	Discount int`json:"discount"`
	Description string`json:"description"`
	Exp string`json:"exp"`
}

type Category struct{
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string
}

type Currency struct{
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string
}

type Unit struct{
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string
}

type Percentage struct{
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Percentage int
}



type MinStock struct{
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Minstock int
	Ppn int
}

type Material struct{
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string
	Satuan int
}

type Customers struct {
	Id     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email  string             `json:"email"`
	Name   string             `json:"name"`
	Adress string             `json:"adress"`
	Phone  string             `json:"phone"`
}

type Carts struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Invoice     string             `json:"invoice"`
	Product_id  primitive.ObjectID `json:"product_id"`
	Qty int `json:"qty"`
	User_id primitive.ObjectID `json:"user_id"`
	// Customer_id primitive.ObjectID `json:"customer_id"`
	// Price int `json:"price"`
	Ppn int`json:"ppn"`
	Percentage int `json:"percentage"`
	Disc int `json:"discount"`
	Price int `json:"price"`
	Total       int               `json:"total"`
	Created_At  time.Time             `json:"created_at"`
	Updated_At  time.Time             `json:"updated_at"`
}

type Orders struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	// Cart_id    int                 `json:"cart_id"`
	// Qty        int                `json:"qty"`
	Price      int                `json:"price"`
	Created_At string             `json:"created_at"`
}

type Profile struct{
	Name string `json:name`
	Phone string `json:phone`
	Zip int `json:zip`
	Description string `json:description`
	Address string `json:address`
}

type Role struct{
	Id         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string `json:name`
}