package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function to collect wikipedia page text into an array
func mapPage(pId uint64, title string, text string, word string) map[string]interface{} {
	// var clc []string
	// for _, title := range pages {
	// 	pId := title.Id
	// 	pageTxt := getPageText(pId)
	// 	clc = append(clc, pageTxt)
	// }
	// return clc
	m := make(map[string]interface{})

	m["pid"] = pId
	m["title"] = title
	m["text"] = text
	m["count"] = countWord(text, word)
	return m
}

func countWord(text string, word string) int {
	var cnt int
	for _, term := range strings.Fields(text) {
		if strings.ToLower(term) == strings.ToLower(word) {
			cnt += 1
		}
	}
	return cnt
}

func insertDb(m map[string]interface{}) {
	mongoURI := "mongodb://root:password@localhost:27017/"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	defer client.Disconnect(ctx)
	fmt.Println(result)
}
