package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"
)

func (ph *AuthHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	addressID, err := strconv.Atoi(r.URL.Query().Get("address_id"))
	if err != nil {
		http.Error(w, "Invalid address_id", http.StatusBadRequest)
		return
	}

	newAddress := r.URL.Query().Get("address")

	updatedAddress, err := ph.Service.AuthService.UpdateAddress(id, addressID, newAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Address updated successfully", updatedAddress)
}
