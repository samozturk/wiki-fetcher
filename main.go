package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type rawResult struct {
	BatchComplete string          `json:"batchcomplete"`
	Continue      string          `json:"continue"`
	Query         randQueryResult `json:"query"`
}

// Structs for getting random page titles query
type randQueryResult struct {
	Pages []randRes `json:"random"`
}

type randRes struct {
	Id string `json:"id"`
	// Ns    int    `json:"ns"`
	Title string `json:"title"`
}

// Structs for page query
type queryResult struct {
	Pages []pageRes `json:"random"`
}

type pageRes struct {
	PageId  string `json:"id"`
	Ns      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

var randomResult rawResult

func main() {

	myRes := getRandomTitles(10)
	fmt.Println(myRes[1].Title)
}

// Function to get random titles
func getRandomTitles(n int) []randRes {
	var randomResult rawResult
	// Format query string
	queryString := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=random&rnnamespace=0&rnlimit=%v&format=json", n)
	// Get response
	resp, err := http.Get(queryString)
	if err != nil {
		log.Fatalln(err)
	}
	// Read body as bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Unmarshall it into an instance of a struct if json structure is valid
	checkValid := json.Valid(body)
	if checkValid {
		json.Unmarshal(body, &randomResult)
	}
	return randomResult.Query.Pages
}

// Function to get page text
func getPageText(tit string) pageRes {
	var pageResult pageRes
	queryString := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&prop=extracts&titles=%v&explaintext=True&format=json", tit)
	// Get response
	resp, err := http.Get(queryString)
	if err != nil {
		log.Fatalln(err)
	}
	// Read body as bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Unmarshall it into an instance of a struct if json structure is valid
	checkValid := json.Valid(body)
	if checkValid {
		json.Unmarshal(body, &pageResult)
	}
	return pageResult
}
