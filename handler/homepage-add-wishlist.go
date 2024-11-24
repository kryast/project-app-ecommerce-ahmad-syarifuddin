package handler

import (
	"fmt"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ph *ProductHandler) AddToWishlistHandler(w http.ResponseWriter, r *http.Request) {

	id := helper.GetID(w, r)
	wishlist, err := ph.Service.ProductService.AddToWishlist(id)
	if err != nil {
		ph.Logger.Error("Error adding to wishlist", zap.Error(err))
		http.Error(w, fmt.Sprintf("Error adding to wishlist: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Success", wishlist)
}
