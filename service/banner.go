package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllBanner() ([]model.Banner, error) {
	banners, err := ps.Repo.ProductRepo.GetAllBanner()
	if err != nil {
		ps.Repo.ProductRepo.Logger.Error("Error fetching banners", zap.Error(err))
		return nil, err
	}

	return banners, nil
}
