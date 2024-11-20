package repository

import (
	"fmt"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) DeleteFromWishlist(wishlistID int) error {
	query := `
		DELETE FROM wishlists 
		WHERE product_details_id = $1;
	`

	_, err := repo.DB.Exec(query, wishlistID)
	if err != nil {
		repo.Logger.Error("Error deleting from wishlist", zap.Error(err))
		return fmt.Errorf("failed to delete from wishlist: %w", err)
	}

	return nil
}
