package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (r *AuthRepositoryDB) GetProfileByID(id int) (*model.Profile, error) {
	query := `
		SELECT id, name, email, phone
		FROM auths 
		WHERE id = $1;
	`

	var profile model.Profile
	err := r.DB.QueryRow(query, id).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Warn("User not found", zap.Int("userID", id), zap.Error(err))
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		r.Logger.Error("Failed to fetch user profile", zap.Int("userID", id), zap.Error(err))
		return nil, err
	}

	return &profile, nil
}
