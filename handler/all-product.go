package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"go.uber.org/zap"
)

type ProductHandler struct {
	Service service.AllService
	Logger  *zap.Logger
	Config  util.Configuration
}

func NewProductHandler(service service.AllService, logger *zap.Logger, config util.Configuration) ProductHandler {
	return ProductHandler{
		Service: service,
		Logger:  logger,
		Config:  config,
	}
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.Service.ProductService.GetAllProducts()
	if err != nil {
		ph.Logger.Error("Error getting products", zap.Error(err))
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", products)
}
