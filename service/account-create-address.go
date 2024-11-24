package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (as *AuthService) AddAddress(authID int, address string, flagDefault bool) (*model.Address, error) {
	newAddress, err := as.Repo.AuthRepo.AddAddress(authID, address, flagDefault)
	if err != nil {
		as.Logger.Error("Failed to add address", zap.Int("authID", authID), zap.String("address", address), zap.Error(err))
		return nil, err
	}

	return newAddress, nil
}
