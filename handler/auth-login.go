package handler

import (
	"encoding/json"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (lh *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login model.Login

	// Decode the JSON body from the request
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the user credentials
	user, err := lh.Service.AuthService.Login(&login)
	if err != nil {
		lh.Logger.Error("Failed to login", zap.Error(err))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate a unique session token
	token := uuid.New().String()

	// Store the token in the database, associated with the user (auth_id)
	err = lh.Service.AuthService.StoreSessionToken(user.ID, token)
	if err != nil {
		lh.Logger.Error("Failed to store token in database", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send success response with user data and token
	helper.SuccessResponse(w, 200, "Login successful", map[string]interface{}{
		"user":  user,
		"token": token, // Return the token in the response, since it's stored in the DB
	})
}
