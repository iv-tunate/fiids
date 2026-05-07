package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string 	`json:"name"`
	Email     string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func UserDTO(dbUser database.User) User{
	return  User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		Email: dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
	}
}