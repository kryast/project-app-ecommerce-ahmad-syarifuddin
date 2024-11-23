package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (repo *AuthRepositoryDB) AddAddress(authID int, address string, flagDefault bool) (*model.Address, error) {

	var newID int
	err := repo.DB.QueryRow(`
		INSERT INTO address (auth_id, address, flag_default)
		VALUES ($1, $2, $3)
		RETURNING id
	`, authID, address, flagDefault).Scan(&newID)

	if err != nil {
		return nil, fmt.Errorf("failed to add address: %v", err)
	}

	// Return the newly added address with ID and flag_default
	return &model.Address{
		ID:          newID,
		AuthID:      authID,
		Address:     address,
		FlagDefault: flagDefault,
	}, nil
}
