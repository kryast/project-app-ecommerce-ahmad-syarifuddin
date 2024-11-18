package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int
	Message string
	Data    interface{}
}

func SuccessResponse(w http.ResponseWriter, code int, message string) {
	response := Response{
		Status:  code,
		Message: message,
		Data:    nil,
	}

	w.WriteHeader(code)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
