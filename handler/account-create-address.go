package handler

import (
	"database/sql"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ah *AuthHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	address := r.URL.Query().Get("address")
	if address == "" {
		ah.Logger.Warn("Address parameter missing", zap.Int("authID", id))
		return
	}

	flagDefault := false
	var existingID int
	err := ah.Service.AuthService.Repo.AuthRepo.DB.QueryRow(`
		SELECT id FROM address WHERE auth_id = $1 AND flag_default = TRUE LIMIT 1
	`, id).Scan(&existingID)

	if err == sql.ErrNoRows {
		flagDefault = true
		ah.Logger.Info("No default address found, setting this as default", zap.Int("authID", id))
	} else if err == nil {
		flagDefault = false
		ah.Logger.Info("Existing default address found, setting new address as non-default", zap.Int("authID", id))
	} else {
		ah.Logger.Error("Error checking default address", zap.Int("authID", id), zap.Error(err))
		return
	}

	newAddress, err := ah.Service.AuthService.AddAddress(id, address, flagDefault)
	if err != nil {
		ah.Logger.Error("Failed to add address", zap.Int("authID", id), zap.String("address", address), zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Address added successfully", newAddress)
}
