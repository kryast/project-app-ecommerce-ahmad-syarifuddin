package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (ph *ProductHandler) UpdateCartQuantity(w http.ResponseWriter, r *http.Request) {

	id := helper.GetID(w, r)

	quantityParam := chi.URLParam(r, "quantity")

	newQuantity, err := strconv.Atoi(quantityParam)
	if err != nil || newQuantity <= 0 {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	err = ph.Service.ProductService.UpdateCartQuantity(id, newQuantity)
	if err != nil {
		http.Error(w, "Failed to update cart quantity", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", nil)
}
