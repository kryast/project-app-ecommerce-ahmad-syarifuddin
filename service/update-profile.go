package service

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (s *AuthService) UpdateProfile(id int, profile *model.Profile) (*model.Profile, error) {
	return s.Repo.AuthRepo.UpdateProfile(id, profile)
}
