package utils

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/lib/pq"
)

func ParseDbError(err error) (int, string) {
	if err == nil {
		return http.StatusOK, "Operation Successful"
	}

	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound, "The requested record could not be found."
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr){
		switch pqErr.Code{
		case "23503":
			return  http.StatusNotFound , "The requested resource or dependency does not exist."

		case "23505": 
			return http.StatusConflict, "This record or relationship already exists."

		case "23514": 
			return http.StatusBadRequest, "The provided data violates system validation rules."
            
		case "22001":
			return http.StatusBadRequest, "One or more text fields exceed the allowed character limit."
		}
	}

	return http.StatusInternalServerError, "An unexpected database error occurred."
}