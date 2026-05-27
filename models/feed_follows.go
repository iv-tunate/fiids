package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/database"
)

type FeedFollow struct {
	ID        uuid.UUID   `json:"id"`
	FeedId    uuid.UUID   `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID   `json:"user_id"`
}

func FeedFollowDTO(feedFollow database.FeedFollow) FeedFollow{
	return FeedFollow{
		ID: feedFollow.ID,
		FeedId: feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID: feedFollow.UserID,
	}
}

func FollowedFeedsDTO(followedFeeds []database.FeedFollow) []FeedFollow{
	followed_feeds := []FeedFollow{}

	for _, ff := range followedFeeds{
		followed_feeds = append(followed_feeds, FeedFollowDTO(ff))
	}

	return  followed_feeds
}