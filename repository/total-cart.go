package repository

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
		return 0, err
	}

	return totalQuantity, nil
}
