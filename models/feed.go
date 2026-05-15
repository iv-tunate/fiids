package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/database"
)

type Feed struct {
	ID        uuid.UUID  `json:"id"`
	Name      string    `json:"name"`
	Url      string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID uuid.UUID	`json:"owner_id"`
}

func FeedDTO(dbFeed database.Feed) Feed{
	return Feed{
		ID: dbFeed.ID,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID: dbFeed.UserID,
	}
}

func FeedsDTO(dbfeeds []database.Feed) []Feed{
	feeds := []Feed{}

	for _, feed := range dbfeeds{
		feeds = append(feeds, FeedDTO(feed))
	}
	return feeds
}