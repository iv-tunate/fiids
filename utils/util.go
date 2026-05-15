package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
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
	Page int
}

func NewPagination(w http.ResponseWriter, r *http.Request) Pagination{
	p := r.URL.Query().Get("page")
	if p == "" {
		p = "1"
	}

	ps := r.URL.Query().Get("page_size")
	if ps == "" {
		ps = "10"
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("Error parsing page parameter '%s': %v", p, err)
		ErrorResponse(w, 400, "Invalid page parameter", "Bad Request")
		return Pagination{}
	}

	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		log.Printf("Error parsing page_size parameter '%s': %v", ps, err)
		ErrorResponse(w, 400, "Invalid page_size parameter", "Bad Request")
		return Pagination{}
	}

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
		Page: page,
	}
}
