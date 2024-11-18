package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/repository"

	"go.uber.org/zap"
)

type AllService struct {
	SampelService SampelService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		SampelService: NewSampelService(repo, log),
	}
}
