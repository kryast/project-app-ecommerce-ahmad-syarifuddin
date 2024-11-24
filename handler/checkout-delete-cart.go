package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ph *ProductHandler) DeleteCartItemHandler(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)
	err := ph.Service.ProductService.DeleteCartItem(id)
	if err != nil {
		ph.Logger.Error("Failed to delete cart item", zap.Error(err))
		http.Error(w, "Failed to delete cart item", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", nil)
}
