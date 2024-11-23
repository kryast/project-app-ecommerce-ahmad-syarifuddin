package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (as *AuthService) AddAddress(authID int, address string, flagDefault bool) (*model.Address, error) {
	newAddress, err := as.Repo.AuthRepo.AddAddress(authID, address, flagDefault)
	if err != nil {
		return nil, err
	}

	return newAddress, nil
}
