package service

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (ps *ProductService) AddToWishlist(productDetailsID int) (*model.Wishlist, error) {
	wishlist, err := ps.Repo.ProductRepo.AddToWishlist(productDetailsID)
	if err != nil {
		ps.Logger.Error("Failed to add product to wishlist", zap.Int("productDetailsID", productDetailsID), zap.Error(err))
		return nil, fmt.Errorf("failed to add to wishlist: %w", err)
	}
	return wishlist, nil
}
