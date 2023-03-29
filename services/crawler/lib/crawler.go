package lib

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Crawl() {
	fmt.Println("Running crawler")

	c := colly.NewCollector(
	// Configure Domains where to get data
	//colly.AllowedDomains("https://www.bild.de/"),
	)

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		fmt.Printf("Link found: %q -> %s\n\n", element.Text, link)
		c.Visit(element.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	c.OnResponse(func(request *colly.Response) {
		fmt.Println("Status: ", request.StatusCode)
	})

	c.OnError(func(request *colly.Response, err error) {
		fmt.Println("Request URL:", request.Request.URL, "failed with response:", request, "nError:", err)
	})

	c.Visit("https://www.bild.de/")
}
