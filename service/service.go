package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/repository"

	"go.uber.org/zap"
)

type AllService struct {
	AuthService    AuthService
	ProductService ProductService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		AuthService:    NewAuthService(repo, log),
		ProductService: NewProductService(repo, log),
	}
}
