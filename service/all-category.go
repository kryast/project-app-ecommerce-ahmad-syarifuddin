package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllCategory() ([]model.Category, error) {
	// Call the repository to get products
	categories, err := ps.Repo.ProductRepo.GetAllCategory()
	if err != nil {
		// Log and propagate the error
		ps.Repo.ProductRepo.Logger.Error("Error fetching products", zap.Error(err))
		return nil, err
	}

	return categories, nil
}
