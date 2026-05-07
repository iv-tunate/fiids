package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iv-tunate/fiids/database"
	"github.com/iv-tunate/fiids/models"
	"github.com/iv-tunate/fiids/utils"
)

func (cfg *ConfigHandler) RegisterUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Printf("[ERROR] RegisterUser: failed to decode JSON: %v", err)
		utils.ErrorResponse(w, 400, fmt.Sprintf("Error parsing JSON: %v", err), "Bad Request")
		return
	}

	db := cfg.Config.DB
	
	user, err := db.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		Email: params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil{
		log.Printf("[ERROR] RegisterUser: DB error: %v", err)
		utils.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Operation failed: %v", err), "DbUpdateError")
		return
	}
	log.Printf("[INFO] user id:`%v`, name:`%s` successfully created", user.ID, user.Name)
	utils.SuccessResponse(w, 201, models.UserDTO(user), "Operation Successful")
}
