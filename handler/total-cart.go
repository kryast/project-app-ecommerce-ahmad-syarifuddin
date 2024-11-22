package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
)

func (ph *ProductHandler) GetTotalCartQuantity(w http.ResponseWriter, r *http.Request) {
	totalQuantity, err := ph.Service.ProductService.GetTotalCartQuantity()
	if err != nil {
		http.Error(w, "Failed to retrieve cart quantity", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"total_quantity": totalQuantity,
	}

	helper.SuccessResponse(w, 200, "success", response)
}
