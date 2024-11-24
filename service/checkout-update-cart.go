package service

import "go.uber.org/zap"

func (service *ProductService) UpdateCartQuantity(productDetailsID, newQuantity int) error {
	err := service.Repo.ProductRepo.UpdateCartQuantity(productDetailsID, newQuantity)
	if err != nil {
		service.Logger.Error("Failed to update cart quantity", zap.Int("productDetailsID", productDetailsID), zap.Int("newQuantity", newQuantity), zap.Error(err))
		return err
	}
	return nil
}
