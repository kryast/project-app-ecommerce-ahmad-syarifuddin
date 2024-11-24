package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
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
			r.Logger.Error("Address not found for the given auth_id and address_id", zap.Int("authID", authID), zap.Int("addressID", addressID))
			return nil, fmt.Errorf("address not found for auth_id %d and address_id %d", authID, addressID)
		}
		r.Logger.Error("Failed to update address", zap.Int("authID", authID), zap.Int("addressID", addressID), zap.Error(err))
		return nil, fmt.Errorf("failed to update address: %v", err)
	}

	return &updatedAddress, nil
}
