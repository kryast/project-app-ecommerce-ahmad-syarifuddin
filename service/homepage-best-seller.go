package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetBestSellers() ([]*model.ProductBestSeller, error) {
	bestSeller, err := ps.Repo.ProductRepo.GetBestSellers()
	if err != nil {
		ps.Logger.Error("Error fetching bestSeller", zap.Error(err))
		return nil, err
	}

	return bestSeller, nil
}
