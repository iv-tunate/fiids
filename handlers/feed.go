package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/database"
	"github.com/iv-tunate/fiids/middleware"
	"github.com/iv-tunate/fiids/models"
	"github.com/iv-tunate/fiids/utils"
)

func (cfg *ConfigHandler) CreateFeed(rw http.ResponseWriter, r *http.Request){
	type parameter struct{
		Name *string `json:"name"`
		Url *string  `json:"url"`
	}

	userID, ok := r.Context().Value(middleware.UserIdKey).(uuid.UUID)
	if !ok {
		log.Print("[Error] CreateFeed: Invalid userId type in context")
		utils.ErrorResponse(rw, 500, "Internal Server Error", "Internal Server Error")
		return
	}
	
	decoder := json.NewDecoder(r.Body)
	params := parameter{}

	err := decoder.Decode(&params)
		if err != nil{
		log.Printf("[ERROR] CreateFeed: failed to decode JSON: %v", err)
		utils.ErrorResponse(rw, 400, fmt.Sprintf("Invalid or missing required parameters on request body: %v", err), "Bad Request")
		return
	}
		if params.Name == nil || *params.Name == ""{
		log.Printf("[ERROR] CreateFeed: missing required field 'name'")
		utils.ErrorResponse(rw, 400, "Missing required field 'name'", "Bad Request")
		return
	}
	if params.Url == nil || *params.Url == ""{
		log.Printf("[ERROR] CreateFeed: missing required field 'email'")
		utils.ErrorResponse(rw, 400, "Missing required field 'email'", "Bad Request")
		return	
	}

	feed, err := cfg.Config.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name: *params.Name,
		Url: *params.Url,
		UserID: userID,
	})

	if err != nil{
		log.Printf("[Error] CreateFeed: An error occured: %v", err)
		utils.ErrorResponse(rw, 500, "An error occured", "DBUpdate Error")
		return
	}

	log.Printf("[INFO] user with ID:`%v`, created a new feed named:`%s`", userID, feed.Name)
	utils.SuccessResponse(rw, 201, models.FeedDTO(feed), "Operation Successful", nil)
}

func (cfg *ConfigHandler) GetFeeds(rw http.ResponseWriter, r *http.Request){
	
	pagination := utils.NewPagination(rw, r)

	page := int32(pagination.Offset)
	pageSize := int32(pagination.Limit)
	feeds, err := cfg.Config.DB.GetFeeds(r.Context(), database.GetFeedsParams{
		Limit: pageSize,
		Offset: page,
	})
	if err != nil {
		log.Printf("[Error]... An error occured while retrieving all feeds: %v", err)
		utils.ErrorResponse(rw, 500, "An Internal Server error occured", "Internal Server Error")
		return
	}

	log.Print("[Feeds retrieved successfully]")
	utils.SuccessResponse(rw, 200, models.FeedsDTO(feeds), "Operation successful", map[string]any{
		"page": pagination.Page,
		"page_size": pageSize,
		"count": len(models.FeedsDTO(feeds)),
	} )
}