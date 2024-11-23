package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
)

func (ah *AuthHandler) GetAddressesByAuthID(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	addresses, err := ah.Service.AuthService.GetAddressesByAuthID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	helper.SuccessResponse(w, 200, "Success", addresses)
}
