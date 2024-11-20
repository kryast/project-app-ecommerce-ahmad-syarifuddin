package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) GetAllCategory() ([]model.Category, error) {
	query := `
		SELECT id, name from categories
	`

	rows, err := repo.DB.Query(query)
	if err != nil {
		repo.Logger.Error("Error executing query", zap.Error(err))
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var categories []model.Category

	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			repo.Logger.Error("Error scanning row", zap.Error(err))
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		repo.Logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return categories, nil
}
