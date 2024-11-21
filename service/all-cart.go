package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (s *ProductService) GetAllCart() ([]*model.Cart, error) {
	carts, err := s.Repo.ProductRepo.GetAllCart()
	if err != nil {
		s.Repo.ProductRepo.Logger.Error("Error retrieving cart items", zap.Error(err))
		return nil, err
	}
	return carts, nil
}
