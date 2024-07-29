package google_crawler

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly"
)

// CrawlGoogle crawls Google search results for the given terms.
func CrawlGoogle(terms string) (string, error) {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*google.*",
		RandomDelay: 5 * time.Second,
	})

	var results string

	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(terms))
	fmt.Println("Visiting URL:", searchURL)

	c.OnHTML("div.g", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		link, exists := e.DOM.Find("a").Attr("href")
		if title != "" && exists {
			results += fmt.Sprintf("Title: %s\nLink: %s\n\n", title, link)
		}
	})

	err := c.Visit(searchURL)
	if err != nil {
		return "", err
	}

	c.Wait()
	fmt.Println("Crawling complete, results:", results)
	return results, nil
}
