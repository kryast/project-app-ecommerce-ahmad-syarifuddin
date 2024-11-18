package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"project-app-ecommerce-ahmad-syarifuddin/repository"

	"go.uber.org/zap"
)

type AuthService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewAuthService(repo repository.AllRepository, logger *zap.Logger) AuthService {
	return AuthService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s *AuthService) Register(register *model.Register) error {

	err := s.Repo.AuthRepo.Register(register)
	if err != nil {
		s.Repo.AuthRepo.Logger.Error("Failed to create new register", zap.Error(err))
		return fmt.Errorf("failed to create new register: %w", err)
	}

	return nil
}
