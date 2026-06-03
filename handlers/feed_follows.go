package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/database"
	"github.com/iv-tunate/fiids/middleware"
	"github.com/iv-tunate/fiids/models"
	"github.com/iv-tunate/fiids/utils"
)

func (cfg *ConfigHandler) FollowFeed(rw http.ResponseWriter, r *http.Request){
	type parameters struct{
		FeedId *uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("[Error] FollowFeed: Failed to decode json body: %v", err)
		utils.ErrorResponse(rw, 400, fmt.Sprintf("Invalid or missing required parameters on request body: %v", err), "Bad Request")
		return
	}

	if params.FeedId == nil{
		log.Printf("[ERROR] FollowFeed: missing required field 'feed_id'")
		utils.ErrorResponse(rw, 400, "Missing required field 'feed_id'", "Bad Request")
		return
	}

	user_id, ok := r.Context().Value(middleware.UserIdKey).(uuid.UUID)

	if !ok{
		log.Print("\n[Error] FollowFeed: Invalid or missing user ID in context\n")
		utils.ErrorResponse(rw, 500, "Internal Server Error", "Internal Server Error")
		return
	}

	feed_follow, err := cfg.Config.DB.FollowFeeds(r.Context(), database.FollowFeedsParams{
		FeedID: *params.FeedId,
		UserID: user_id,
	})

	if err != nil{
		log.Printf("\n[ERROR] FollowFeed: An error occured while user with ID: %v attempted to follow a feed... Details: %v \n", user_id, err)

		statusCode, msg := utils.ParseDbError(err)
		utils.ErrorResponse(rw, statusCode, msg, http.StatusText(statusCode))	
		return	
	}

	log.Printf("[Success] FollowFeed: Feed followed successfully by user: %v", user_id)
	utils.SuccessResponse(rw, 201, models.FeedFollowDTO(feed_follow), "Operation successful", nil)
}

func (cfg *ConfigHandler) GetFollowedFeeds(rw http.ResponseWriter, r *http.Request){

	userId, ok := r.Context().Value(middleware.UserIdKey).(uuid.UUID)
	if !ok{
		log.Print("\n[Error] FollowFeed: Invalid or missing user ID in context\n")
		utils.ErrorResponse(rw, 500, "Internal Server Error", "Internal Server Error")
		return
	}

	followed_feeds, err := cfg.Config.DB.GetFollowedFeeds(r.Context(), userId)

	if err != nil{
		log.Printf("[ERROR] GetFollowedFeeds: An error occured while fetching User with ID:%v's followed feeds. Details: %v", userId, err)
		statusCode, msg := utils.ParseDbError(err)
		utils.ErrorResponse(rw, statusCode, msg, http.StatusText(statusCode))
		return
	}

	log.Printf("[SUCCESS] GetFollowedFeeds: Followed feeds for user with ID:%v retrieved successfully", userId)
	utils.SuccessResponse(rw, 200, models.FollowedFeedsDTO(followed_feeds), "Operation Successful", nil)
}

func (cfg ConfigHandler) UnfollowFeed(rw http.ResponseWriter, r *http.Request){
	userId, ok := r.Context().Value(middleware.UserIdKey).(uuid.UUID)
	if !ok{
		log.Print("[Error] DeleteFollowedFeed: An error occured while trying to fetch the user id from the context")
		utils.ErrorResponse(rw, http.StatusInternalServerError, "An error occured", http.StatusText(500))
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil{
		log.Printf("[Error] UnfollowFeed: An error occured while parsing retrieved followed feed Id: %v... \nError Details: %v", idStr, err)
		utils.ErrorResponse(rw, 400, "Invalid or broken feed id passed", http.StatusText(400))
		return
	}

	err = cfg.Config.DB.DeleteFollowedFeeds(r.Context(), database.DeleteFollowedFeedsParams{
		ID: id,
		UserID: userId,
	})

	if err != nil{
		log.Printf("[Error] UnfollowFeed: An error occured while trying to execute delete operation on the database for a followed feed wit ID: %v.\n Error Details: %v\n", id, err);
		statusCode, msg := utils.ParseDbError(err)
		utils.ErrorResponse(rw, statusCode, msg, http.StatusText(statusCode))
		return
	}

	log.Println("[Success] UnfollowFeed: Operation Successful.")
	utils.SuccessResponse(rw, 200, nil, "Feed Unfollowed Successfully", nil)
}