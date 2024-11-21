package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"time"

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

func (repo *ProductRepositoryDB) GetAllProducts() ([]*model.Product, error) {
	query := `
		SELECT 
			pd.id,
			p.name AS product_name,
			p.detail AS product_detail,
			p.price AS product_price,
			c.name AS category_name,
			ph.photo_url,
			p.created_at
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

	var products []*model.Product

	for rows.Next() {
		var product model.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Detail, &product.Price, &product.Category, &product.PhotoURL, &product.CreatedAt)
		if err != nil {
			repo.Logger.Error("Error scanning row", zap.Error(err))
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		daysSinceCreation := time.Since(product.CreatedAt).Hours() / 24
		if daysSinceCreation < 30 {
			product.FlagNew = "true"
		} else {
			product.FlagNew = ""
		}

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		repo.Logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return products, nil
}
