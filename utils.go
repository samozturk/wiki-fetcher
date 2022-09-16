package main

import (
	"strings"
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
