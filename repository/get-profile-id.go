package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (r *AuthRepositoryDB) GetProfileByID(id int) (*model.Profile, error) {
	query := `
		SELECT a.id AS id, a.name, a.email, a.phone, b.address
		FROM auths a
		LEFT JOIN address b ON a.id = b.auth_id
		WHERE a.id = $1
	`

	var profile model.Profile
	err := r.DB.QueryRow(query, id).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.Phone, &profile.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	return &profile, nil
}
