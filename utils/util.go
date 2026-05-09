package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func ParseGuidFromHttpReq(r *http.Request) (uuid.UUID, error) {
	guidStr := r.URL.Query().Get("guid")
	if guidStr == ""{
		log.Printf("Invalid guid: %v ", guidStr)
		return uuid.UUID{}, errors.New("Invalid guid parameter")
		}
	return uuid.Parse(guidStr)
}

func GenerateRandomKey(length int) (string, error){
	if length < 1 {
		length = 32
	}

	bytes:= make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil{
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}