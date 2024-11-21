package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status    int
	Message   string
	Page      int         `json:"page,omitempty"`
	TotalPage int         `json:"total_page,omitempty"`
	Data      interface{} `json:"data,omitempty"`
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

func SuccessResponseFilter(w http.ResponseWriter, code int, message string, page int, total_page int, data any) {
	successResponse := Response{
		Status:    code,
		Message:   message,
		Page:      page,
		TotalPage: total_page,
		Data:      data,
	}
	response, _ := json.MarshalIndent(successResponse, "", " ")
	jsonResponse, _ := w.Write(response)
	json.NewEncoder(w).Encode(jsonResponse)
}
