package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MyCollection struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	zipCode    string             `bson:"zipCode"`
	supplierId string             `bson:"supplierId,omitempty"`
	oemNumber  string             `bson:"oemNumber,omitempty" json:"oem_number,omitempty"`
}

func main() {

	uri := "mongodb://127.0.0.1:27117,127.0.0.1:27118"
	fmt.Printf("uri: %v\n", uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	defer client.Disconnect(context.TODO())

	database := client.Database("persons")
	mycol := database.Collection("personscollection", &options.CollectionOptions{})

	for i := 101003; i < 9999900; i++ {

		// obj := &struct{ personid int }{personid: 131005}
		// fmt.Printf("obj: %v\n", obj)

		res, err := mycol.InsertOne(
			context.Background(),
			bson.M{"personid":i/* 131005 */},
		)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("res: %v\n", res)

	}

}

func mainold() {
	uri := "mongodb://127.0.0.1:27117,127.0.0.1:27118"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	defer client.Disconnect(context.TODO())

	database := client.Database("MyDatabase")
	mycol := database.Collection("MyCollection3", &options.CollectionOptions{})

	for i := 0; i < 99999; i++ {

		FF := "partitionnumber1"
		rrr := rand.Float64()
		rand.Seed(time.Now().UnixNano())
		// fmt.Printf("rrr: %v\n", rrr)

		if rrr < 0.5 {
			FF = "partitionNumber2"
		}
		obj := &MyCollection{
			oemNumber:  "3434",
			zipCode:    FF,
			supplierId: "spdocsdomcosd",
		}
		fmt.Printf("obj: %v\n", obj)

		res, err := mycol.InsertOne(
			context.Background(),
			obj,
		)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("res: %v\n", res)

	}

}
