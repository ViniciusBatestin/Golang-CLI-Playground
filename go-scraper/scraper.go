package main

import (
	"fmt"

	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type item struct {
	name     string `json: "name"`
	price    string `json: "price"`
	imageURL string `json: "imgurl"`
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

	var items []item

	c.OnHTML("div.grid-product__content", func(h *colly.HTMLElement) {
		item := item{
			name:     h.ChildText(".grid-product__title.grid-product__title--body"),
			price:    h.ChildText(".grid-product__price"),
			imageURL: extractImageURL(h.ChildAttr(".grid__image-ratio", "style")),
		}
		items = append(items, item)
	})

	c.OnHTML("a.pagination__next", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		if next_page != "" {
			c.Visit(next_page)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	eslLink := ("https://shop.esl.com/collections/jerseys")
	err := c.Visit(eslLink)
	if err != nil {
		fmt.Println("Error visiting the page", err)
	}
	fmt.Println(items)
}
