package helper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetID(w http.ResponseWriter, r *http.Request) int {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		SuccessResponse(w, 400, "Invalid ID", id)
		return 0
	}
	return id
}
