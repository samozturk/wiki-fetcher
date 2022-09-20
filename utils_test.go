package main

import "testing"

func TestCoundWord(t *testing.T) {
	testStr := "this is a Test program to test this program"

	res := countWord(testStr, "test")
	exp := 2
	if res != exp {
		t.Errorf("Expected %d got %d instead", exp, res)

	}
}

func TestMapPage(t *testing.T) {
	exp := make(map[string]interface{})
	exp["pid"] = uint64(1)
	exp["title"] = "Test Title"
	mapTxt := "The quick brown fox jumps over the lazy quick dog"
	exp["text"] = "The quick brown fox jumps over the lazy quick dog"
	exp["count"] = 2

	res := mapPage(1, "Test Title", mapTxt, "quick")
	if res["pid"] != exp["pid"] {
		t.Errorf("Expected %d got %d instead", exp["pid"], res["pid"])
	}
	if res["title"] != exp["title"] {
		t.Errorf("Expected %v got %v instead", exp["title"], res["title"])
	}
	if res["text"] != exp["text"] {
		t.Errorf("Expected %q got %q instead", exp["text"], res["text"])
	}
	if res["count"] != exp["count"] {
		t.Errorf("Expected %d got %d instead", exp["count"], res["count"])
	}

}

func TestInsertDb(t *testing.T) {
	expMap := make(map[string]interface{})
	expMap["first field"] = 1
	expMap["second field"] = "two"
	insertDb(expMap)

}
