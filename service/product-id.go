package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (s *ProductService) GetProductById(productId int) (*model.Product, error) {
	product, err := s.Repo.ProductRepo.GetProductById(productId)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	return product, nil
}
