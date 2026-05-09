package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseDetail struct{
	Code int `json:"statuscode"`
	Data any `json:"data"`
	PaginationDetail map[string]any `json:"pagination_detail"`
	Msg string `json:"message"`
	Error any `json:"error"`
}

func Response(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)

	if err != nil{
		log.Fatalf("Failed to marshal JSON response: %v", payload)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ErrorResponse(w http.ResponseWriter, code int, message string, err any) {
	if code > 499{
		log.Printf("Internal Server Error 5xx: %s", message)
	}
	response := ResponseDetail{
		Code: code,
		Msg: message,
		Error: err,
	}

	Response(w, code, response)
}

func SuccessResponse(w http.ResponseWriter, code int, data any, message string, paginationDetail map[string]any){
	response := ResponseDetail{
		Code: code,
		Data: data,
		PaginationDetail: paginationDetail,
		Msg: message,
	}
	Response(w, code, response)
}