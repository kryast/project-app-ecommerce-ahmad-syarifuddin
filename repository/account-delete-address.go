package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (r *AuthRepositoryDB) DeleteAddress(authID int, addressID int) (*model.Address, error) {
	var deletedAddress model.Address
	query := `
		DELETE FROM address
		WHERE id = $1 AND auth_id = $2
		RETURNING id, auth_id, address, flag_default
	`

	err := r.DB.QueryRow(query, addressID, authID).Scan(&deletedAddress.ID, &deletedAddress.AuthID, &deletedAddress.Address, &deletedAddress.FlagDefault)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Warn("Address not found for deletion", zap.Int("authID", authID), zap.Int("addressID", addressID))
			return nil, fmt.Errorf("address not found for auth_id %d and address_id %d", authID, addressID)
		}
		r.Logger.Error("Failed to delete address", zap.Int("authID", authID), zap.Int("addressID", addressID), zap.Error(err))
		return nil, fmt.Errorf("failed to delete address: %v", err)
	}

	return &deletedAddress, nil
}
