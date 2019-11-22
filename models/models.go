package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID       primitive.ObjectID`json:"_id" bson:"_id,omitempty"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Products struct {
	Id    primitive.ObjectID    `json:"_id" bson:"_id,omitempty"`
	Name  string    `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Customers struct {
	Id     primitive.ObjectID    `json:"_id" bson:"_id,omitempty"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Adress string    `json:"adress"`
	Phone  string    `json:"phone"`
}

type Carts struct {
	Id          primitive.ObjectID	`json:"_id" bson:"_id,omitempty"`
	Invoice     string `json:"invoice"`
	Product_id 	primitive.ObjectID `json:"product_id"`
	Customer_id primitive.ObjectID `json:"customer_id"`
	Total       int    `json:"total"`
	Created_At  string `json:"created_at"`
	Updated_At  string `json:"updated_at"`
}

type Orders struct {
	Id         primitive.ObjectID`json:"_id" bson:"_id,omitempty"`
	Cart_id    int    `json:"cart_id"`
	Qty        int    `json:"qty"`
	Price      int    `json:"price"`
	Created_At string `json:"created_at"`
}
