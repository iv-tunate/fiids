package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/database"
	"github.com/iv-tunate/fiids/models"
	"github.com/iv-tunate/fiids/utils"
)

func (cfg *ConfigHandler) RegisterUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name *string `json:"name"`
		Email *string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("[ERROR] RegisterUser: failed to decode JSON: %v", err)
		utils.ErrorResponse(w, 400, fmt.Sprintf("Error parsing JSON: %v", err), "Bad Request")
		return
	}

	if params.Name == nil || *params.Name == ""{
		log.Printf("[ERROR] RegisterUser: missing required field 'name'")
		utils.ErrorResponse(w, 400, "Missing required field 'name'", "Bad Request")
		return
	}
	if params.Email == nil || *params.Email == ""{
		log.Printf("[ERROR] RegisterUser: missing required field 'email'")
		utils.ErrorResponse(w, 400, "Missing required field 'email'", "Bad Request")
		return	
	}

	db := cfg.Config.DB
	
	user, err := db.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: *params.Name,
		Email: *params.Email,
	})

	if err != nil{
		log.Printf("[ERROR] RegisterUser: DB error: %v", err)
		utils.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Operation failed: %v", err), "DbUpdateError")
		return
	}
	log.Printf("[INFO] user id:`%v`, name:`%s` successfully created", user.ID, user.Name)
	utils.SuccessResponse(w, 201, models.UserDTO(user), "Operation Successful")
}

func (cfg *ConfigHandler) GetUserById(w http.ResponseWriter, r *http.Request){
	userId, err := utils.ParseGuidFromHttpReq(r)
	if err != nil{
		log.Printf("An error due to an invalid guid value occured while fetching user by ID...[ERROR] GetUserById: %v", err)
		utils.ErrorResponse(w, 400, "[Error]Invalid user ID", "Bad Request")
		return
	}

	user, err := cfg.Config.DB.GetUserById(r.Context(), userId)
	if err != nil{
		log.Printf("An error occured while fetching user by ID...[ERROR] GetUserById: %v", err)
		utils.ErrorResponse(w, 404, fmt.Sprintf("[Error]User not found: %v", err), "Not Found")
		return
	}
	log.Printf("[INFO] user id:`%v`, name:`%s` successfully fetched", user.ID, user.Name)
	utils.SuccessResponse(w, 200, models.UserDTO(user), "Operation Successful")
}