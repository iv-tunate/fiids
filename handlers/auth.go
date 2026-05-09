package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/iv-tunate/fiids/database"
	"github.com/iv-tunate/fiids/utils"
	response "github.com/iv-tunate/fiids/utils"
)

func (cfg *ConfigHandler) GenerateApiKey(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name *string `json:"name"`
	}
	userId, err := utils.ParseGuidFromHttpReq(r)

	if err != nil{
		response.ErrorResponse(w, 400, "Invalid user ID", "Bad Request")
		return
	}
	
	_, err = cfg.Config.DB.GetUserById(r.Context(), userId)
	if err != nil{
		log.Printf("...Invalid user... [ERROR] GenerateApiKey: DB error: %v", err)
		response.ErrorResponse(w, 404, "User not found", "Not Found")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}

	err = decoder.Decode(params)
	if err != nil || params.Name == nil{
		log.Printf("...Invalid request body... [ERROR] GenerateApiKey: %v", err)
		response.ErrorResponse(w, 400, "Invalid request body or name parameter", "Bad Request")
		return
	}
	randomKey, err := utils.GenerateRandomKey(32)
	if err != nil{
		log.Printf("...Failed to generate API key... [ERROR] GenerateApiKey: %v", err)
		response.ErrorResponse(w, 500, "Failed to generate API key", "Internal Server Error")
		return
	}
	apiKey := "Fiids_" + strings.TrimSpace(*params.Name) + "_" + randomKey

	hashedKey := sha256.Sum256([]byte(apiKey))
	encodedHash :=hex.EncodeToString(hashedKey[:])

	_, err = cfg.Config.DB.GenerateApiKey(r.Context(), database.GenerateApiKeyParams{
		Name: *params.Name,
		ApiKey: encodedHash,
		UserID: userId,
	})
	if err != nil{
		log.Printf("...Failed to generate API key... [ERROR] GenerateApiKey: %v", err)
		response.ErrorResponse(w, 500, "Failed to generate API key", "Internal Server Error")
		return
	}
	response.SuccessResponse(w, 200, map[string]string{"api_key": apiKey}, "API key generated successfully")
}