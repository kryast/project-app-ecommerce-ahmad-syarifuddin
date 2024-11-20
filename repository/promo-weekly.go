package repository

import "project-app-ecommerce-ahmad-syarifuddin/model"

func (r *ProductRepositoryDB) GetPromoWeekly() ([]*model.PromoWeekly, error) {
	query := `
         SELECT 
            id, 
            product_details_id, 
            TO_CHAR(start_week, 'YYYY-MM-DD') AS start_week,
            TO_CHAR(end_week, 'YYYY-MM-DD') AS end_week
        FROM promo_weekly
        WHERE NOW()::DATE BETWEEN start_week AND end_week;
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promotions []*model.PromoWeekly
	for rows.Next() {
		var promotion model.PromoWeekly
		if err := rows.Scan(&promotion.ID, &promotion.ProductDetailsID, &promotion.StartWeek, &promotion.EndWeek); err != nil {
			return nil, err
		}
		promotions = append(promotions, &promotion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return promotions, nil
}
