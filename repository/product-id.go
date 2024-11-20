package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (repo *ProductRepositoryDB) GetProductById(productId int) (*model.Product, error) {
	// Query untuk mengambil data produk berdasarkan id
	query := `
		SELECT 
			pd.id AS id,
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
		WHERE 
			pd.id = $1
	`

	rows := repo.DB.QueryRow(query, productId)

	var product model.Product

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Detail,
		&product.Price,
		&product.Category,
		&product.PhotoURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item with id %d not found", productId)
		}
		return nil, err
	}

	return &product, nil
}
