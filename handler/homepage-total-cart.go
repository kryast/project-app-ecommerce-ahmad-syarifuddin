package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetTotalCartQuantity(w http.ResponseWriter, r *http.Request) {
	totalQuantity, err := ph.Service.ProductService.GetTotalCartQuantity()
	if err != nil {
		ph.Logger.Error("Failed to fetch total cart quantity", zap.Error(err))
		return
	}

	response := map[string]interface{}{
		"total_quantity": totalQuantity,
	}

	helper.SuccessResponse(w, 200, "success", response)
}
