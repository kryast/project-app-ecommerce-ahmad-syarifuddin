package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ph *AuthHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	existingProfile, err := ph.Service.AuthService.GetProfileByID(id)
	if err != nil {
		ph.Logger.Warn("Profile not found", zap.Int("userID", id), zap.Error(err))
		return
	}

	if name == "" {
		name = *existingProfile.Name
	}
	if email == "" {
		email = *existingProfile.Email
	}

	profile := model.Profile{
		ID:    id,
		Name:  &name,
		Email: &email,
	}

	updatedProfile, err := ph.Service.AuthService.UpdateProfile(id, &profile)
	if err != nil {
		ph.Logger.Error("Failed to update profile", zap.Int("userID", id), zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Profile updated successfully", updatedProfile)
}
