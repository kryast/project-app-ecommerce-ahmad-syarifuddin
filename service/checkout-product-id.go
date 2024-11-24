package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetProductById(productId int) (*model.Product, error) {
	product, err := ps.Repo.ProductRepo.GetProductById(productId)
	if err != nil {
		ps.Logger.Error("Failed to get product by ID", zap.Int("productId", productId), zap.Error(err))
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	return product, nil
}
