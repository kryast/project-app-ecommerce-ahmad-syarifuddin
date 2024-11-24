package repository

import "go.uber.org/zap"

func (repo *ProductRepositoryDB) GetTotalCartQuantity() (int, error) {
	query := `
        SELECT 
            SUM(quantity) AS total_quantity
        FROM 
            transaction_items
        WHERE 
            transaction_id = 1;
    `

	var totalQuantity int
	err := repo.DB.QueryRow(query).Scan(&totalQuantity)
	if err != nil {
		repo.Logger.Error("Failed to fetch total cart quantity", zap.Error(err))
		return 0, err
	}

	return totalQuantity, nil
}
