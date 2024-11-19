package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

type ProductRepositoryDB struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewProductRepository(db *sql.DB, Log *zap.Logger) ProductRepositoryDB {
	return ProductRepositoryDB{
		DB:     db,
		Logger: Log,
	}
}

func (repo *ProductRepositoryDB) GetAllProducts() ([]model.Product, error) {
	query := `
		SELECT 
			p.name AS product_name,
			p.detail AS product_detail,
			p.price AS product_price,
			c.name AS category_name,
			ph.photo_url
		FROM 
			product_details pd
		JOIN 
			products p ON p.id = pd.product_id
		JOIN 
			categories c ON c.id = pd.category_id
		JOIN 
			photos ph ON ph.id = pd.photo_id
	`

	rows, err := repo.DB.Query(query)
	if err != nil {
		repo.Logger.Error("Error executing query", zap.Error(err))
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []model.Product

	// Iterate over the rows and scan into the Product struct
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Name, &product.Detail, &product.Price, &product.Category, &product.PhotoURL)
		if err != nil {
			repo.Logger.Error("Error scanning row", zap.Error(err))
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		products = append(products, product)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		repo.Logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return products, nil
}
