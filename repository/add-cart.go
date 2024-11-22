package repository

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) AddToCart(productDetailsID int, quantity int) error {
	var existingQuantity int
	err := repo.DB.QueryRow(`
		SELECT quantity
		FROM transaction_items
		WHERE transaction_id = 1 AND product_details_id = $1
	`, productDetailsID).Scan(&existingQuantity)

	if err != nil && err != sql.ErrNoRows {
		repo.Logger.Error("Error checking product in cart", zap.Error(err), zap.Int("product_details_id", productDetailsID))
		return fmt.Errorf("failed to check product in cart: %w", err)
	}

	if err == sql.ErrNoRows {
		query := `
			INSERT INTO transaction_items (transaction_id, product_details_id, quantity)
			VALUES (1, $1, $2);
		`
		_, err := repo.DB.Exec(query, productDetailsID, quantity)
		if err != nil {
			repo.Logger.Error("Error adding to cart", zap.Error(err), zap.Int("product_details_id", productDetailsID), zap.Int("quantity", quantity))
			return fmt.Errorf("failed to add to cart: %w", err)
		}
	} else {
		query := `
			UPDATE transaction_items
			SET quantity = quantity + $1
			WHERE transaction_id = 1 AND product_details_id = $2;
		`
		_, err := repo.DB.Exec(query, quantity, productDetailsID)
		if err != nil {
			repo.Logger.Error("Error updating quantity in cart", zap.Error(err), zap.Int("product_details_id", productDetailsID), zap.Int("quantity", quantity))
			return fmt.Errorf("failed to update quantity in cart: %w", err)
		}
	}

	return nil
}
