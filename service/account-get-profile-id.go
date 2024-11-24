package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (as *AuthService) GetProfileByID(id int) (*model.Profile, error) {
	profile, err := as.Repo.AuthRepo.GetProfileByID(id)
	if err != nil {
		as.Logger.Error("Failed to fetch profile", zap.Int("userID", id), zap.Error(err))
		return nil, fmt.Errorf("profile not found for ID %d", id)
	}
	return profile, nil
}
