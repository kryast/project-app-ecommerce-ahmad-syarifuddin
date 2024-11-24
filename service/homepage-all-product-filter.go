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

func (s *ProductService) GetAllProductsFilter(nameTerm, categoryTerm string, page, limit int) ([]*model.Product, int, error) {
	return s.Repo.ProductRepo.GetAllProductsFilter(nameTerm, categoryTerm, page, limit)
}
