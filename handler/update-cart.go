package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"
)

func (ph *ProductHandler) UpdateCartQuantity(w http.ResponseWriter, r *http.Request) {

	productDetailsIDStr := r.URL.Query().Get("product_details_id")
	newQuantityStr := r.URL.Query().Get("new_quantity")

	productDetailsID, err := strconv.Atoi(productDetailsIDStr)
	if err != nil {
		http.Error(w, "Invalid product_details_id", http.StatusBadRequest)
		return
	}

	newQuantity, err := strconv.Atoi(newQuantityStr)
	if err != nil {
		http.Error(w, "Invalid new_quantity", http.StatusBadRequest)
		return
	}

	err = ph.Service.ProductService.UpdateCartQuantity(productDetailsID, newQuantity)
	if err != nil {
		http.Error(w, "Failed to update cart quantity", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", nil)
}
