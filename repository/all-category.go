package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) GetAllCategory() ([]model.Category, error) {
	query := `
		SELECT name from categories
	`

	rows, err := repo.DB.Query(query)
	if err != nil {
		repo.Logger.Error("Error executing query", zap.Error(err))
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var categories []model.Category

	// Iterate over the rows and scan into the Product struct
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Name)
		if err != nil {
			repo.Logger.Error("Error scanning row", zap.Error(err))
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, category)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		repo.Logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return categories, nil
}
