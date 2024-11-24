package service

import "go.uber.org/zap"

func (ps *ProductService) GetTotalCartQuantity() (int, error) {
	totalQuantity, err := ps.Repo.ProductRepo.GetTotalCartQuantity()
	if err != nil {
		ps.Logger.Error("Failed to fetch total cart quantity", zap.Error(err))
		return 0, err
	}
	return totalQuantity, nil
}
