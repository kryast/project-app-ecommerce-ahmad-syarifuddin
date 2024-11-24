package repository

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (r *ProductRepositoryDB) GetAllRecomment() ([]*model.Recomment, error) {
	query := `
        SELECT 
            r.id,
            p.name AS title,
            p.detail AS subtitle,
            ph.photo_url AS photo, 
            r.product_details_id AS product_id
        FROM 
            recomments r
        JOIN 
            product_details pd ON pd.id = r.product_details_id
        JOIN 
            products p ON p.id = pd.product_id
        JOIN 
            photos ph ON ph.id = pd.photo_id;
    `

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recomments []*model.Recomment
	for rows.Next() {
		var recomment model.Recomment
		if err := rows.Scan(&recomment.ID, &recomment.Product, &recomment.ProductDetail, &recomment.PhotoURL, &recomment.ProductID); err != nil {
			return nil, err
		}
		recomments = append(recomments, &recomment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recomments, nil
}
