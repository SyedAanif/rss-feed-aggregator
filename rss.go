package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// get marshalled XML RSS feed from a URL
func urlToFeed(url string) (RSSFeed, error){
	// Set up an HTTP client to get RSS feed
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// Make GET request to RSS url
	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	
	// defer to close response body till return (finally)
	defer resp.Body.Close()

	// Read all response body to byte arrays(slice)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	// empty struct to hold values
	rssFeed := RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}