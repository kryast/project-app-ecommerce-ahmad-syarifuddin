package service

import (
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllRecomment() ([]*model.Recomment, error) {
	recomments, err := ps.Repo.ProductRepo.GetAllRecomment()
	if err != nil {
		ps.Logger.Error("Error retrieving recommendations", zap.Error(err))
		return nil, err
	}
	return recomments, nil
}
