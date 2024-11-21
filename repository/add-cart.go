package repository

import (
	"fmt"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) AddToCart(productDetailsID int, quantity int) error {
	query := `
        INSERT INTO transaction_items (transaction_id, product_details_id, quantity)
        VALUES (1, $1, $2);`

	_, err := repo.DB.Exec(query, productDetailsID, quantity)
	if err != nil {
		repo.Logger.Error("Error adding to cart", zap.Error(err), zap.Int("product_details_id", productDetailsID), zap.Int("quantity", quantity))
		return fmt.Errorf("failed to add to cart: %w", err)
	}

	return nil
}
