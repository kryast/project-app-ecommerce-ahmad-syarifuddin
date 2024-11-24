package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (r *ProductRepositoryDB) GetAllBanner() ([]*model.Banner, error) {

	query := `
		SELECT 
			b.id,
			ph.photo_url,
			b.title,
			b.subtitle,
			b.path_page
		FROM 
			banners b
		JOIN 
			photos ph ON ph.id = b.photo_id  
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		r.Logger.Error("Error executing query", zap.Error(err))
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var banners []*model.Banner

	for rows.Next() {
		var banner model.Banner
		err := rows.Scan(&banner.ID, &banner.PhotoURL, &banner.Title, &banner.Subtitle, &banner.PathPage)
		if err != nil {
			r.Logger.Error("Error scanning row", zap.Error(err))
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		banners = append(banners, &banner)
	}

	if err := rows.Err(); err != nil {
		r.Logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return banners, nil
}
