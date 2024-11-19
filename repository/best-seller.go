package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) GetBestSellers() ([]model.ProductBestSeller, error) {
	query := `
        SELECT 
            p.name AS product_name,
            SUM(ti.quantity) AS total_sold
        FROM 
            transaction_items ti
        JOIN 
            product_details pd ON pd.id = ti.product_details_id
        JOIN 
            products p ON p.id = pd.product_id
        JOIN 
            transactions t ON t.id = ti.transaction_id
        WHERE 
            t.id = 2 
        GROUP BY 
			p.name
        ORDER BY 
            total_sold DESC
        LIMIT 5;
    `

	rows, err := repo.DB.Query(query)
	if err != nil {
		repo.Logger.Error("Error executing query", zap.Error(err))
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var bestSellers []model.ProductBestSeller

	for rows.Next() {
		var bestSeller model.ProductBestSeller
		err := rows.Scan(&bestSeller.Name, &bestSeller.TotalSold)
		if err != nil {
			repo.Logger.Error("Error scanning row", zap.Error(err))
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		bestSellers = append(bestSellers, bestSeller)
	}

	if err := rows.Err(); err != nil {
		repo.Logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return bestSellers, nil
}
