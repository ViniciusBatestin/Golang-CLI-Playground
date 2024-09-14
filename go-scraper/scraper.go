package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	imageURL string `json: "imgurl"`
	name     string `json: "name"`
	price    string `json: "price"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("shop.esl.com"),
	)

	c.OnHTML("div.grid-product__content", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText(".grid-product__title.grid-product__title--body"))
	})

	eslLink := ("https://shop.esl.com/?ref=ratcomp&utm_campaign=gs-2020-04-15&utm_source=google&utm_medium=smart_campaign&gad_source=1&gclid=CjwKCAjw6JS3BhBAEiwAO9waF_q7P91RVVwGHOdHSRwxBS-2LFB7nI2d60P338Sy8o-sYKv59Voh_hoCBIYQAvD_BwE")
	err := c.Visit(eslLink)
	if err != nil {
		fmt.Println("Error visiting the page", err)
	}
}
