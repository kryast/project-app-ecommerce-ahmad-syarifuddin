package repository

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (repo *ProductRepositoryDB) GetAllCart() ([]*model.Cart, error) {
	query := `
        SELECT 
            ti.id,
            p.name AS product_name,
            p.price AS product_price,
            SUM(ti.quantity) AS total_quantity,
            SUM(ti.quantity * p.price) AS product_subtotal
        FROM 
            transaction_items ti
        JOIN 
            product_details pd ON pd.id = ti.product_details_id
        JOIN 
            products p ON p.id = pd.product_id
        JOIN 
            transactions t ON t.id = ti.transaction_id
        WHERE 
            t.id = 1 
        GROUP BY 
            ti.id, p.name, p.price
        ORDER BY 
            total_quantity DESC
        LIMIT 5;
    `

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []*model.Cart
	// var totalPrice float64

	for rows.Next() {
		var cart model.Cart
		if err := rows.Scan(&cart.ID, &cart.Product, &cart.Price, &cart.Quantity, &cart.Subtotal); err != nil {
			return nil, err
		}
		// cart.TotalPrice = cart.Subtotal
		// totalPrice += cart.Subtotal
		carts = append(carts, &cart)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// for i := range carts {
	// 	carts[i].TotalPrice = totalPrice
	// }

	return carts, nil
}
