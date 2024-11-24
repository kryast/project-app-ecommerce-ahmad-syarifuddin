package service

import (
	"fmt"

	"go.uber.org/zap"
)

func (ps *ProductService) DeleteCartItem(productDetailsID int) error {
	err := ps.Repo.ProductRepo.DeleteFromCart(productDetailsID)
	if err != nil {
		ps.Logger.Error("Failed to delete product from cart", zap.Int("productDetailsID", productDetailsID), zap.Error(err))
		return fmt.Errorf("failed to delete product from cart: %w", err)
	}

	return nil
}
