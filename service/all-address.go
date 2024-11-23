package service

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (as *AuthService) GetAddressesByAuthID(authID int) ([]model.Address, error) {
	addresses, err := as.Repo.AuthRepo.GetAddressesByAuthID(authID)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
