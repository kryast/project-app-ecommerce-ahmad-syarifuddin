package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"

	"go.uber.org/zap"
)

func (ph *AuthHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	addressID, err := strconv.Atoi(r.URL.Query().Get("address_id"))
	if err != nil {
		ph.Logger.Error("Invalid address_id provided", zap.String("address_id", r.URL.Query().Get("address_id")), zap.Error(err))
		return
	}

	newAddress := r.URL.Query().Get("address")
	if newAddress == "" {
		ph.Logger.Error("Missing new address", zap.Int("authID", id))
		return
	}

	updatedAddress, err := ph.Service.AuthService.UpdateAddress(id, addressID, newAddress)
	if err != nil {
		ph.Logger.Error("Failed to update address", zap.Int("authID", id), zap.Int("addressID", addressID), zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Address updated successfully", updatedAddress)
}
