package handler

import (
	"encoding/json"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (lh *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login model.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := lh.Service.AuthService.Login(&login)
	if err != nil {
		lh.Logger.Error("Failed to login", zap.Error(err))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	helper.SuccessResponse(w, 200, "Login successful", user)
}
