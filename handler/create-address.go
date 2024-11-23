package handler

import (
	"database/sql"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
)

func (ah *AuthHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	flagDefault := false
	var existingID int
	err := ah.Service.AuthService.Repo.AuthRepo.DB.QueryRow(`
		SELECT id FROM address WHERE auth_id = $1 AND flag_default = TRUE LIMIT 1
	`, id).Scan(&existingID)

	if err == sql.ErrNoRows {
		flagDefault = true
	} else if err == nil {
		flagDefault = false
	} else {
		http.Error(w, "Error checking default address", http.StatusInternalServerError)
		return
	}

	newAddress, err := ah.Service.AuthService.AddAddress(id, address, flagDefault)
	if err != nil {
		http.Error(w, "Failed to add address: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Address added successfully", newAddress)
}
