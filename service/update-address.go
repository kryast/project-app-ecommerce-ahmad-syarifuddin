package service

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (s *AuthService) UpdateAddress(authID int, addressID int, newAddress string) (*model.Address, error) {
	updatedAddress, err := s.Repo.AuthRepo.UpdateAddress(authID, addressID, newAddress)
	if err != nil {
		return nil, err
	}

	return updatedAddress, nil
}
