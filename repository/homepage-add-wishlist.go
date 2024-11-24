package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) AddToWishlist(productDetailsID int) (*model.Wishlist, error) {
	query := `
		INSERT INTO wishlists (product_details_id)
		VALUES ($1) RETURNING id, product_details_id;
	`

	var wishlist model.Wishlist
	err := repo.DB.QueryRow(query, productDetailsID).Scan(&wishlist.ID, &wishlist.ProductDetailsID)
	if err != nil {
		repo.Logger.Error("Error adding to wishlist", zap.Error(err))
		return nil, fmt.Errorf("failed to add to wishlist: %w", err)
	}

	return &wishlist, nil
}
