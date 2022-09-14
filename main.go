package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	rand_Queries := getRandomTitles(10)
	wikis := collectTxt(rand_Queries)
	fmt.Println(wikis[9])

	// wikiMap = make(map[string]interface{})
	for _, j := range rand_Queries {
		fmt.Println(j.Id)
		fmt.Println(j.Title)
		fmt.Println("---")
	}

}

func insertDb() {
	mongoURI := "mongodb://root:password@localhost:27017/"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)

	demoDB := client.Database("freenow")
	// defer catsCollection.Drop(ctx)

	err = demoDB.CreateCollection(ctx, "taxi")
	if err != nil {
		log.Fatal(err)
	}
	taxiCollection := demoDB.Collection("taxi")

	result, err := taxiCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Mocha"},
		{Key: "breed", Value: "Turkish Van"},
	})
	fmt.Println(result)
}

func init() {
	// // Create db connection
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Ping db to check connection
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// collection = client.Database("tasker").Collection("tasks")
	// log.Println("DONE")
}
