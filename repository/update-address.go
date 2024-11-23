package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (r *AuthRepositoryDB) UpdateAddress(authID int, addressID int, newAddress string) (*model.Address, error) {
	var updatedAddress model.Address
	query := `
		UPDATE address
		SET address = $1
		WHERE id = $2 AND auth_id = $3
		RETURNING id, auth_id, address
	`

	err := r.DB.QueryRow(query, newAddress, addressID, authID).Scan(&updatedAddress.ID, &updatedAddress.AuthID, &updatedAddress.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("address not found for auth_id %d and address_id %d", authID, addressID)
		}
		return nil, fmt.Errorf("failed to update address: %v", err)
	}

	return &updatedAddress, nil
}
