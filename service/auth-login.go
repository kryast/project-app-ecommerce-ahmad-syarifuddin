package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (as *AuthService) Login(login *model.Login) (*model.Register, error) {
	user, err := as.Repo.AuthRepo.Login(login)
	if err != nil {
		as.Logger.Error("Failed to login user", zap.Error(err))
		return nil, fmt.Errorf("failed to login user: %w", err)
	}

	return user, nil
}
