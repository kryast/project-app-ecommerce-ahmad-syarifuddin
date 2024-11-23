package service

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (s *AuthService) DeleteAddress(authID int, addressID int) (*model.Address, error) {
	deletedAddress, err := s.Repo.AuthRepo.DeleteAddress(authID, addressID)
	if err != nil {
		return nil, err
	}

	return deletedAddress, nil
}
