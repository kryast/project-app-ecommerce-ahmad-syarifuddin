package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ah *AuthHandler) GetAddressesByAuthID(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	addresses, err := ah.Service.AuthService.GetAddressesByAuthID(id)
	if err != nil {
		ah.Logger.Error("Failed to fetch addresses", zap.Int("authID", id), zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Success", addresses)
}
