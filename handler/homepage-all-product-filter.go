package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"
	"strconv"

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

func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	nameTerm := r.URL.Query().Get("name")
	categoryTerm := r.URL.Query().Get("category")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 3

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 3
		}
	}

	products, total, err := h.Service.ProductService.GetAllProductsFilter(nameTerm, categoryTerm, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPages := (total + limit - 1) / limit
	helper.SuccessResponseFilter(w, 200, "Success", page, totalPages, products)

}
