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

func SuccessResponse(w http.ResponseWriter, code int, message string, data any) {
	successResponse := Response{
		Status:  code,
		Message: message,
		Data:    data,
	}
	response, _ := json.MarshalIndent(successResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}
