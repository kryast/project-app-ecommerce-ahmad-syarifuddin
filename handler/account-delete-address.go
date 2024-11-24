package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"

	"go.uber.org/zap"
)

func (ph *AuthHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	addressID, err := strconv.Atoi(r.URL.Query().Get("address_id"))
	if err != nil {
		ph.Logger.Error("Invalid address_id", zap.String("addressID", r.URL.Query().Get("address_id")), zap.Error(err))
		return
	}

	deletedAddress, err := ph.Service.AuthService.DeleteAddress(id, addressID)
	if err != nil {
		ph.Logger.Error("Failed to delete address", zap.Int("authID", id), zap.Int("addressID", addressID), zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Address deleted successfully", deletedAddress)
}
