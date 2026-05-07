package handlers

import (
	
	"net/http"
	response "github.com/iv-tunate/fiids/utils"
)


func HandlerHealth(w http.ResponseWriter, r *http.Request){
	response.SuccessResponse(w, 200, map[string]string{"status": "ok"}, "operation successful")
}

func HandlerError(w http.ResponseWriter, r *http.Request){
	response.ErrorResponse(w, 400, "Something went wrong", "Bad Request")	
}