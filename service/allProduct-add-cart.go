package service

import "go.uber.org/zap"

func (s *ProductService) AddToCart(id int, quantity int) error {
	err := s.Repo.ProductRepo.AddToCart(id, quantity)
	if err != nil {
		s.Logger.Error("Failed to add product to cart", zap.Int("productId", id), zap.Int("quantity", quantity), zap.Error(err))
		return err
	}
	return nil
}
