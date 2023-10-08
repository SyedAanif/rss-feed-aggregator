package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/SyedAanif/rss-feed-aggregator/internal/database"
)

/*
	This function will run periodically in the background to scrape
	RSS Feed on given time interval durations.
*/
func startScrapping(
	db *database.Queries, // connection to DB
	concurrency int, // how many go routines for scrapping
	timeBetweenRequest time.Duration, // request time interval for scrapping
){
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	// To keep track of passed duration it passes a tick on a CHANNEL
	ticker := time.NewTicker(timeBetweenRequest)

	// Initialised like this to get the first tick immediately and then wait for tick duration
	for ; ; <-ticker.C{
		// Get batch of feeds based on concurrency
		// from a global context of GO application
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			log.Println("Error fetching feeds:",err)
			continue // always running function
		}

		// A WaitGroup waits for a collection of goroutines to finish. 
		// The main goroutine calls Add to set the number of goroutines to wait for. 
		// Then each of the goroutines runs and calls Done when finished. 
		// At the same time, Wait can be used to block until all goroutines have finished.
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1) // add a go routiine for fetching, this will be equivalent to concurrency

			go scrapeFeed(db, feed, wg) // scrape feed on go routine
		}
		wg.Wait() // blocking operation to wait for all go routines to finish
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup){
	defer wg.Done() // will defer done or decrement of each routine once function returns

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID) // mark the feed as fetched
	if err != nil {
		log.Println("Error marking feed as fetched:",err)
		return
	}

	// Get actual feed for the URL
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed:",err)
		return
	}

	// Log to console
	for _ , item := range rssFeed.Channel.Item{
		log.Printf("Found post: %v on feed name: %v",item.Title, feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}