package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) CreateOrder(authID int) (*model.Order, error) {
	order, err := ps.Repo.ProductRepo.GetOrderData(authID)
	if err != nil {
		ps.Logger.Error("Failed to fetch order data", zap.Int("authID", authID), zap.Error(err))
		return nil, fmt.Errorf("failed to fetch order data: %w", err)
	}

	return order, nil
}
