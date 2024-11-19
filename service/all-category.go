package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllCategory() ([]model.Category, error) {
	categories, err := ps.Repo.ProductRepo.GetAllCategory()
	if err != nil {
		ps.Repo.ProductRepo.Logger.Error("Error fetching products", zap.Error(err))
		return nil, err
	}

	return categories, nil
}
