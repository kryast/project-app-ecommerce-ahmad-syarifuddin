package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (as *AuthService) GetProfileByID(id int) (*model.Profile, error) {
	profile, err := as.Repo.AuthRepo.GetProfileByID(id)
	if err != nil {
		return nil, fmt.Errorf("profile not found for ID %d", id)
	}
	return profile, nil
}
