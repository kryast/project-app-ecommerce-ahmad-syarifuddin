package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (service *ProductService) CreateOrder(authID int) (*model.Order, error) {
	order, err := service.Repo.ProductRepo.GetOrderData(authID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order data: %w", err)
	}

	return order, nil
}
