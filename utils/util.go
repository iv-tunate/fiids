package utils

import (
	"crypto/rand"
	"encoding/hex"
)

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

type Pagination struct{
	Offset int
	Limit int
}
func NewPagination(page int, pageSize int) Pagination{
	if page < 1{
		page = 1
	}
	if pageSize < 1{
		pageSize =10
	}
	if pageSize > 100{
		pageSize = 100
	}

	return Pagination{
		Limit: pageSize,
		Offset: (page - 1) * pageSize,
	}
}
