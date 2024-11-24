package handler

import (
	"fmt"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (ph *ProductHandler) AddToCart(w http.ResponseWriter, r *http.Request) {

	id := helper.GetID(w, r)

	quantityParam := chi.URLParam(r, "quantity")
	quantity := 1

	if quantityParam != "" {
		parsedQuantity, err := strconv.Atoi(quantityParam)
		if err != nil || parsedQuantity <= 0 {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}
		quantity = parsedQuantity
	}

	err := ph.Service.ProductService.AddToCart(id, quantity)
	if err != nil {
		ph.Logger.Error("Error adding to cart", zap.Error(err))
		http.Error(w, fmt.Sprintf("Error adding to cart: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Success", nil)
}
