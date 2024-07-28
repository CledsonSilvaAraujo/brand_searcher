package tests

import (
	"backend/google_crawler"
	"testing"
)

func TestCrawlGoogle(t *testing.T) {
	results, err := google_crawler.CrawlGoogle("example terms")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if results == "" {
		t.Fatalf("Expected results, got empty string")
	}
}
