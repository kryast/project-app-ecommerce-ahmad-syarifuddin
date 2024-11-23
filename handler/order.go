package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
)

func (ph *ProductHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)
	order, err := ph.Service.ProductService.CreateOrder(id)
	if err != nil {
		helper.SuccessResponse(w, 404, "ID not Found", nil)
		return
	}

	helper.SuccessResponse(w, 200, "success", order)
}
