package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	product, err := ph.Service.ProductService.GetProductById(id)
	if err != nil {
		ph.Logger.Error("Error fetching product by ID", zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Success", product)
}
