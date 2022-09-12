package main

import "strings"

// Function to collect wikipedia page text into an array
func collectTxt(pages []randRes) []string {
	var clc []string
	for _, title := range pages {
		pId := title.Id
		pageTxt := getPageText(pId)
		clc = append(clc, pageTxt)
	}
	return clc
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
