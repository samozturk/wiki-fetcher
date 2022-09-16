package main

// Structs for random query //
type randRawResult struct {
	BatchComplete string          `json:"batchcomplete"`
	Continue      string          `json:"continue"`
	Query         randQueryResult `json:"query"`
}
type randQueryResult struct {
	Pages []randRes `json:"random"`
}
type randRes struct {
	Id uint64 `json:"id"`
	// Ns    int    `json:"ns"`
	Title string `json:"title"`
}

// Stuct for database
type pageWordCount struct {
	pId       int
	text      string
	wordCount int
}
