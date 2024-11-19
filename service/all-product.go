package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"project-app-ecommerce-ahmad-syarifuddin/repository"

	"go.uber.org/zap"
)

type ProductService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewProductService(repo repository.AllRepository, logger *zap.Logger) ProductService {
	return ProductService{
		Repo:   repo,
		Logger: logger,
	}
}

func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := ps.Repo.ProductRepo.GetAllProducts()
	if err != nil {
		ps.Repo.ProductRepo.Logger.Error("Error fetching products", zap.Error(err))
		return nil, err
	}

	return products, nil
}
