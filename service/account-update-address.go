package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (as *AuthService) UpdateAddress(authID int, addressID int, newAddress string) (*model.Address, error) {
	updatedAddress, err := as.Repo.AuthRepo.UpdateAddress(authID, addressID, newAddress)
	if err != nil {
		as.Logger.Error("Failed to update address", zap.Int("authID", authID), zap.Int("addressID", addressID), zap.Error(err))
		return nil, err
	}

	return updatedAddress, nil
}
