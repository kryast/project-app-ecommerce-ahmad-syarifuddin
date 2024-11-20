package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (s *ProductService) GetAllRecomment() ([]*model.Recomment, error) {
	recomments, err := s.Repo.ProductRepo.GetAllRecomment()
	if err != nil {
		s.Repo.ProductRepo.Logger.Error("Error retrieving recommendations", zap.Error(err))
		return nil, err
	}
	return recomments, nil
}
