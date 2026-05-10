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
}

func FeedDTO(dbUser database.Feed) Feed{
	return Feed{
		ID: dbUser.UserID,
		Name: dbUser.Name,
		Url: dbUser.Url,
		CreatedAt: dbUser.CreatedAt,
	}
}