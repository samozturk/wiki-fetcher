package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	for i := 0; i <= 9; i++ {
		pId, title := getRandomTitleId()
		text := getPageText(pId)
		m := mapPage(pId, title, text, "and")
		insertDb(m)
	}
	os.Exit(0)
}

func insertDb(m map[string]interface{}) {
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

	// err = demoDB.CreateCollection(ctx, "taxi")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	taxiCollection := demoDB.Collection("taxi")

	result, err := taxiCollection.InsertOne(ctx, m)
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
