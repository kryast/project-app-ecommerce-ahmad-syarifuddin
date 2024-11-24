package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetPromoWeekly() ([]*model.PromoWeekly, error) {
	promotions, err := ps.Repo.ProductRepo.GetPromoWeekly()
	if err != nil {
		ps.Logger.Error("Failed to get weekly promotions", zap.Error(err))
		return nil, err
	}
	return promotions, nil
}
