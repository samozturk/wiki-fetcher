package main

import (
	"fmt"
)

func main() {
	// rand_Queries := getRandomTitles(10)
	// wikis := collectTxt(rand_Queries)
	// fmt.Println(wikis[9])
	myStr := "hello there Hello my name is sam and I am coding go"
	cnt := countWord(myStr, "hello")
	fmt.Println(cnt)

}
