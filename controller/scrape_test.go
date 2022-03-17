package controller

import (
	"testing"
)


func TestScrape(t *testing.T) {
	m, err := Scrape()

	if err != nil{
		t.Fatal(err)
	}

	t.Log(string(m))
} 
