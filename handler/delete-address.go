package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"
)

func (ph *AuthHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	addressID, err := strconv.Atoi(r.URL.Query().Get("address_id"))
	if err != nil {
		http.Error(w, "Invalid address_id", http.StatusBadRequest)
		return
	}

	deletedAddress, err := ph.Service.AuthService.DeleteAddress(id, addressID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Address deleted successfully", deletedAddress)
}
