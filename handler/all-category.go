package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	// Fetch products using the service layer
	categories, err := ph.Service.ProductService.GetAllCategory()
	if err != nil {
		ph.Logger.Error("Error getting products", zap.Error(err))
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", categories)
}
