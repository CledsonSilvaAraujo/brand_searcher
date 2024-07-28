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
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*google.*",
		RandomDelay: 5 * time.Second,
	})

	var results string

	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(terms))

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.ChildAttr("a", "href")
		results += fmt.Sprintf("Title: %s\nLink: %s\n\n", title, link)
	})

	err := c.Visit(searchURL)
	if err != nil {
		return "", err
	}

	c.Wait()
	return results, nil
}
