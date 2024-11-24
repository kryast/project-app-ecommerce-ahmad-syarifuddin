package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (as *AuthService) DeleteAddress(authID int, addressID int) (*model.Address, error) {
	deletedAddress, err := as.Repo.AuthRepo.DeleteAddress(authID, addressID)
	if err != nil {
		as.Logger.Error("Failed to delete address", zap.Int("authID", authID), zap.Int("addressID", addressID), zap.Error(err))
		return nil, err
	}

	return deletedAddress, nil
}
