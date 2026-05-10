package middleware

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/iv-tunate/fiids/config"
	"github.com/iv-tunate/fiids/utils"
)

type contextKey string
const UserIdKey contextKey = "userID"

func ApiKeyAuthMiddleware(cfg *config.ApiConfig) func(http.Handler) http.Handler{
	result := func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(rw http.ResponseWriter, r * http.Request){
			apiKey, err := extractApiKey(r.Header)
			if err != nil{
				log.Printf("...[Error]... An error occured: %v", err)
				utils.ErrorResponse(rw, 401, err.Error(), "Unauthorized")
				return 
			}

				hashedKey := sha256.Sum256([]byte(apiKey))
				encodedHash := hex.EncodeToString(hashedKey[:])

				dbHash, err := cfg.DB.CheckApiKey(r.Context(), encodedHash)
				if err != nil{
					if errors.Is(err, sql.ErrNoRows){
						log.Printf("...[Invalid ApiKey Error]... %v", err)
						utils.ErrorResponse(rw, 401, "Invalid ApiKey", "Unauthorized")
						return 
					}
					log.Printf("...[Error]... An error occured: %v", err)
					utils.ErrorResponse(rw, 500, "An error occured", "Internal Server Error")
					return 
				}

				if dbHash.RevokedAt.Valid{
					log.Printf("...[Error]... Access attempted by user with ID: %v with revoked apikey %v at %v", dbHash.UserID, dbHash.ApiKey, time.Now())
					utils.ErrorResponse(rw, 403, "ApiKey access revoked", "Forbidden")
					return 
				}
				
				ctx := context.WithValue(r.Context(), UserIdKey, dbHash.UserID)
				r = r.WithContext(ctx)
				next.ServeHTTP(rw, r)
		})
	}

	return result
}
func extractApiKey(h http.Header) (string, error){
	headerVal := h.Get("Authorization")
	if headerVal == ""{
		return "", errors.New("Authorization header is empty")
	}

	apiKey := strings.Split(headerVal, " ")
	if len(apiKey) != 2 && apiKey[0] != "ApiKey"{
		return "", errors.New("Malformed or invalid auth header")
	}

	return  apiKey[1], nil
}