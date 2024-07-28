package google_crawler

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

// CrawlGoogle crawls Google search results for the given terms.
func CrawlGoogle(terms string) (string, error) {
	c := colly.NewCollector()

	var results string

	encodedTerms := url.QueryEscape(terms)
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", encodedTerms)
	log.Printf("Visiting URL: %s", searchURL)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/url?q=") {
			actualLink := strings.Split(link, "&")[0]
			actualLink = strings.TrimPrefix(actualLink, "/url?q=")
			results += actualLink + "\n"
		}
	})

	// Handle request errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Printf("Error visiting URL: %v", err)
	})

	err := c.Visit(searchURL)
	if err != nil {
		log.Printf("Error visiting URL: %v", err)
		return "", err
	}

	return results, nil
}
