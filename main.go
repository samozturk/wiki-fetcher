package main

import "os"

func main() {
	for i := 0; i <= 9; i++ {
		pId, title := getRandomTitleId()
		text := getPageText(pId)
		m := mapPage(pId, title, text, "and")
		insertDb(m)
	}
	os.Exit(0)

}
