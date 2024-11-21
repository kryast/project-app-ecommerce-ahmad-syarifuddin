package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"time"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) GetAllProductsFilter(nameTerm, categoryTerm string, page, limit int) ([]*model.Product, int, error) {
	// Base query with filtering
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
        WHERE 
            ($1::text IS NULL OR p.name ILIKE $1) AND
            ($2::text IS NULL OR c.name ILIKE $2)
        LIMIT $3 OFFSET $4
    `

	// Prepare arguments for the query
	var args []interface{}
	if nameTerm != "" {
		args = append(args, fmt.Sprintf("%%%s%%", nameTerm))
	} else {
		args = append(args, nil) // Use NULL for no filtering
	}
	if categoryTerm != "" {
		args = append(args, fmt.Sprintf("%%%s%%", categoryTerm))
	} else {
		args = append(args, nil) // Use NULL for no filtering
	}

	// Calculate offset
	offset := (page - 1) * limit
	args = append(args, limit, offset)

	// Fetch products
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		repo.Logger.Error("Error executing query", zap.Error(err))
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Detail, &product.Price, &product.Category, &product.PhotoURL, &product.CreatedAt); err != nil {
			repo.Logger.Error("Error scanning row", zap.Error(err))
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}

		daysSinceCreation := time.Since(product.CreatedAt).Hours() / 24
		if daysSinceCreation < 30 {
			product.FlagNew = "true"
		} else {
			product.FlagNew = ""
		}

		products = append(products, &product)
	}

	// Count total products for pagination
	countQuery := `
        SELECT COUNT(*)
        FROM product_details pd
        JOIN products p ON p.id = pd.product_id
        JOIN categories c ON c.id = pd.category_id
        WHERE ($1::text IS NULL OR p.name ILIKE $1) AND
              ($2::text IS NULL OR c.name ILIKE $2)
    `

	var total int
	err = repo.DB.QueryRow(countQuery, args[0], args[1]).Scan(&total)
	if err != nil {
		repo.Logger.Error("Error counting products", zap.Error(err))
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	return products, total, nil
}
