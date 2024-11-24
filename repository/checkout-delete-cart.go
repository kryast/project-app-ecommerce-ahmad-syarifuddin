package repository

import (
	"fmt"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) DeleteFromCart(productDetailsID int) error {

	query := `
		DELETE FROM transaction_items 
		WHERE transaction_id = 1 AND product_details_id = $1;
	`
	_, err := repo.DB.Exec(query, productDetailsID)
	if err != nil {
		repo.Logger.Error("Error deleting product from cart", zap.Error(err), zap.Int("product_details_id", productDetailsID))
		return fmt.Errorf("failed to delete product from cart: %w", err)
	}

	return nil
}
