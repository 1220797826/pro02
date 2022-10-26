package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMonGoDB() {
	co := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error

	client, err = mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success connect")
}

func TestBson() {
	d := bson.D{{"name", "tom"}}
	fmt.Printf("d: %v\n", d)
}

type Student struct {
	Name string
	Age  int
}

func insertOne() {
	s := Student{Name: "tom", Age: 20}
	c2 := client.Database("golang_db").Collection("student")
	// defer c2.Close()
	ior, err := c2.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
}

func insertMany() {
	s1 := Student{Name: "zhang.yu", Age: 22}
	s2 := Student{Name: "jack", Age: 24}
	s := []interface{}{s1, s2}
	c2 := client.Database("golang_db").Collection("student")
	imr, err := c2.InsertMany(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("imr.InsertedIDs: %v\n", imr.InsertedIDs)
}

func update() {
	c2 := client.Database("golang_db").Collection("student")
	ctx := context.TODO()
	update := bson.D{{"$set", bson.D{{"name", "big tom"}, {"age", 99}}}}
	ur, err := c2.UpdateMany(ctx, bson.D{{"name", "tom"}}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.MatchedCount: %v\n", ur.MatchedCount)
}
func find() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c2 := client.Database("golang_db").Collection("student")
	cur, err := c2.Find(ctx, bson.D{{"name", "zhang.yu"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["name"])
	}
}

func del() {
	c2 := client.Database("golang_db").Collection("student")
	ctx := context.TODO()
	dr, _ := c2.DeleteMany(ctx, bson.D{{"name", "zhang.yu"}})
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}
func main() {
	InitMonGoDB()
	del()
}
