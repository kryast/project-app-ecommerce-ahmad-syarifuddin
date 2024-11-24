package repository

import "go.uber.org/zap"

func (repo *ProductRepositoryDB) UpdateCartQuantity(productDetailsID int, newQuantity int) error {
	query := `
        UPDATE transaction_items
        SET quantity = $1
        WHERE transaction_id = 1 AND product_details_id = $2;
    `

	_, err := repo.DB.Exec(query, newQuantity, productDetailsID)
	if err != nil {
		repo.Logger.Error("Failed to update cart quantity", zap.Int("productDetailsID", productDetailsID), zap.Int("newQuantity", newQuantity), zap.Error(err))
		return err
	}

	return nil
}
