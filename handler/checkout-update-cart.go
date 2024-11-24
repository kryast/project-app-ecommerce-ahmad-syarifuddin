package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (ph *ProductHandler) UpdateCartQuantity(w http.ResponseWriter, r *http.Request) {

	id := helper.GetID(w, r)

	quantityParam := chi.URLParam(r, "quantity")

	newQuantity, err := strconv.Atoi(quantityParam)
	if err != nil || newQuantity <= 0 {
		ph.Logger.Error("Invalid quantity received", zap.String("quantityParam", quantityParam), zap.Error(err))
		return
	}

	err = ph.Service.ProductService.UpdateCartQuantity(id, newQuantity)
	if err != nil {
		ph.Logger.Error("Failed to update cart quantity", zap.Int("userId", id), zap.Int("newQuantity", newQuantity), zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "success", nil)
}
