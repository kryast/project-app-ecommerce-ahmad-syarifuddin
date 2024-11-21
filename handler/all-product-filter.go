package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"strconv"
)

func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	// Get the search terms from the query parameters
	nameTerm := r.URL.Query().Get("name")
	categoryTerm := r.URL.Query().Get("category")

	// Get pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Set default values for page and limit
	page := 1
	limit := 3 // Default limit

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

	// Call the service to get products
	products, total, err := h.Service.ProductService.GetAllProductsFilter(nameTerm, categoryTerm, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Calculate total pages
	totalPages := (total + limit - 1) / limit // Ceiling division

	// Set the response header for JSON
	w.Header().Set("Content-Type", "application/json")

	// Respond with the product list and total pages
	// response := struct {
	// 	Products    []*model.Product `json:"products"`
	// 	Total       int              `json:"total"`
	// 	TotalPages  int              `json:"total_pages"`
	// 	CurrentPage int              `json:"current_page"`
	// }{
	// 	Products:    products,
	// 	Total:       total,
	// 	TotalPages:  totalPages,
	// 	CurrentPage: page,
	// }

	helper.SuccessResponseFilter(w, 200, "Success", page, totalPages, products)

}
