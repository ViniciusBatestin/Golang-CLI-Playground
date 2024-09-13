package main

import (
	"github.com/gocolly/colly"
)

type item struct {
	imageURL string `json: "imgurl"`
	name     string `json: "name"`
	price    string `json: "price"`
}

func main() {
	c := colly.NewCollector()
}
