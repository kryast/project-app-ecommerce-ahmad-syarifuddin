package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetBestSellers(w http.ResponseWriter, r *http.Request) {
	bestSeller, err := ph.Service.ProductService.GetBestSellers()
	if err != nil {
		ph.Logger.Error("Error getting bestSeller", zap.Error(err))
		http.Error(w, "Failed to retrieve bestSeller", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", bestSeller)
}
