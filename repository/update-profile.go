package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (r *AuthRepositoryDB) UpdateProfile(id int, profile *model.Profile) (*model.Profile, error) {
	query := `
		UPDATE auths
		SET name = $1, email = $2, phone = $3
		WHERE id = $4
		RETURNING id, name, email, phone
	`
	var updatedProfile model.Profile

	err := r.DB.QueryRow(query, profile.Name, profile.Email, profile.Phone, id).
		Scan(&updatedProfile.ID, &updatedProfile.Name, &updatedProfile.Email, &updatedProfile.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ID %d not found", id)
		}
		return nil, err
	}

	return &updatedProfile, nil
}
