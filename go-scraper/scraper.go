package main

import (
	"fmt"

	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type item struct {
	imageURL string `json: "imgurl"`
	name     string `json: "name"`
	price    string `json: "price"`
}

func extractImageURL(style string) string {
	// Use a regular expression to find the URL inside the background-image property
	re := regexp.MustCompile(`url\((.*?)\)`)
	match := re.FindStringSubmatch(style)
	if len(match) > 1 {
		// Ensure the URL is complete (prefix with https if needed)
		url := match[1]
		if strings.HasPrefix(url, "//") {
			url = "https:" + url
		}
		return url
	}
	return ""
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("shop.esl.com"),
	)

	c.OnHTML("div.grid-product__content", func(h *colly.HTMLElement) {
		item := item{
			name:     h.ChildText(".grid-product__title.grid-product__title--body"),
			price:    h.ChildText(".grid-product__price"),
			imageURL: extractImageURL(h.ChildAttr(".grid__image-ratio", "style")),
		}
		fmt.Println(item)
	})

	eslLink := ("https://shop.esl.com/?ref=ratcomp&utm_campaign=gs-2020-04-15&utm_source=google&utm_medium=smart_campaign&gad_source=1&gclid=CjwKCAjw6JS3BhBAEiwAO9waF_q7P91RVVwGHOdHSRwxBS-2LFB7nI2d60P338Sy8o-sYKv59Voh_hoCBIYQAvD_BwE")
	err := c.Visit(eslLink)
	if err != nil {
		fmt.Println("Error visiting the page", err)
	}
}
