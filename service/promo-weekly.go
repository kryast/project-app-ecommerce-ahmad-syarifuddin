package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (s *ProductService) GetPromoWeekly() ([]*model.PromoWeekly, error) {
	promotions, err := s.Repo.ProductRepo.GetPromoWeekly()
	if err != nil {
		s.Logger.Error("Failed to get weekly promotions", zap.Error(err))
		return nil, err
	}
	return promotions, nil
}
