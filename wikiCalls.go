package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jeffail/gabs"
)

// Function to get random titles
func getRandomTitles(n int) []randRes {
	var randomResult randRawResult
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
func getPageText(pId int) string {

	queryString := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&prop=extracts&pageids=%v&explaintext=True&format=json", pId)

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
	// Parse body
	parsedBody, err := gabs.ParseJSON(body)

	if err != nil {
		log.Fatalln(err)
	}
	// Prep variables for extracting page text
	var value string
	var ok bool
	parsingStr := fmt.Sprintf("query.pages.%v.extract", pId)
	value, ok = parsedBody.Path(parsingStr).Data().(string)

	if ok {
		return value
	} else {
		return "nil"
	}
}
