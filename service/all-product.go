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

// NewProductService creates a new instance of ProductService
func NewProductService(repo repository.AllRepository, logger *zap.Logger) ProductService {
	return ProductService{
		Repo:   repo,
		Logger: logger,
	}
}

// GetProducts retrieves a list of products using the repository
func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
	// Call the repository to get products
	products, err := ps.Repo.ProductRepo.GetAllProducts()
	if err != nil {
		// Log and propagate the error
		ps.Repo.ProductRepo.Logger.Error("Error fetching products", zap.Error(err))
		return nil, err
	}

	return products, nil
}
