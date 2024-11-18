package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"project-app-ecommerce-ahmad-syarifuddin/repository"

	"go.uber.org/zap"
)

type RegisterService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewRegisterService(repo repository.AllRepository, logger *zap.Logger) RegisterService {
	return RegisterService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *RegisterService) Register(register *model.Register) error {

	err := s.Repo.RegisterRepo.Register(register)
	if err != nil {
		s.Repo.RegisterRepo.Logger.Error("Failed to create new register", zap.Error(err))
		return fmt.Errorf("failed to create new register: %w", err)
	}

	return nil
}
