package services

import (
	"context"
	"database/sql"
	"log"
	"strings"
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
		// log.Println("[Success] Found Post...", item.Title, "On feed", feed.Name)
		description := sql.NullString{}
		if item.Description != ""{
			description.String = item.Description
			description.Valid = true
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil{
			log.Printf("Couldn't parse date %v with err %v", item.PubDate, err)
			continue
		}
		_, err = db.CreatePost(ctx, database.CreatePostParams{
			Title: item.Title,
			Description: description,
			PublishedAt: pubDate,
			Url: item.Link,
			FeedID: feed.ID,
		})

		if err != nil{

			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			_, msg := utils.ParseDbError(err)
			log.Printf("[Error] scrapeFeed: An error occured while trying create a post for with feedId: %v.\n Error Details: %v\n", feed.ID, msg);
		}
	}
	log.Printf("[INFO] scrapeFeed: Feed %s collected... %v posts found", feed.Name, len(rssFeed.Channel.Item))
}