package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (r *AuthRepositoryDB) UpdateAddress(authID int, addressID int, newAddress string, flagDefault bool) (*model.Address, error) {
	var updatedAddress model.Address
	query := `
		UPDATE address
		SET address = $1, flag_default = $2
		WHERE id = $3 AND auth_id = $4
		RETURNING id, auth_id, address, flag_default
	`

	err := r.DB.QueryRow(query, newAddress, flagDefault, addressID, authID).Scan(&updatedAddress.ID, &updatedAddress.AuthID, &updatedAddress.Address, &updatedAddress.FlagDefault)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("address not found for auth_id %d and address_id %d", authID, addressID)
		}
		return nil, fmt.Errorf("failed to update address: %v", err)
	}

	return &updatedAddress, nil
}
