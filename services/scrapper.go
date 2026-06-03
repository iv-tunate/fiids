package services

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/iv-tunate/fiids/database"
	"github.com/iv-tunate/fiids/utils"
)

func ScrapeFeeds(ctx context.Context, db *database.Queries, concurrency int, timeBtwnReq time.Duration){
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBtwnReq)

	ticker := time.NewTicker(timeBtwnReq)

	for{
		select{
		case <-ctx.Done():
			log.Printf("Scraper shutting down gracefully at %v\n", time.DateTime)
			return
		case <-ticker.C:
			feeds, err := db.GetNextFeedsToFetch(ctx, int32(concurrency))
			if err != nil{
				_, msg := utils.ParseDbError(err)
				log.Printf("\n[ERROR] Scrape: An error occured while retrieving next feeds to fetch from the database\nError Details: %v\n", msg)
				continue
			}

			wg := &sync.WaitGroup{}
			for _, feed := range feeds{
				wg.Add(1)

				go scrapeFeed(ctx,db, wg, feed)
			}
			wg.Wait()
		}
	}
}

func scrapeFeed(ctx context.Context, db *database.Queries, wg *sync.WaitGroup, feed database.Feed){
	defer wg.Done()
	
	 _, err := db.MarkFeedAsFetched(ctx, feed.ID)
	 if err != nil{
		_, msg := utils.ParseDbError(err)
		log.Printf("\n[ERROR] Scrape: An error occured while retrieving next feeds to fetch from the database\nError Details: %v\n", msg)
		return
	}

	rssFeed, err := utils.UrlToFeed(feed.Url);
	if err != nil{
		log.Printf("[Error] scrapeFeed: An error occured.\nError Details: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item{
		log.Println("[Success] Found Post...", item.Title, "On feed", feed.Name)
	}
	log.Printf("[INFO] scrapeFeed: Feed %s collected... %v posts found", feed.Name, len(rssFeed.Channel.Item))
}