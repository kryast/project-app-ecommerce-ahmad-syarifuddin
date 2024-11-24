package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (repo *AuthRepositoryDB) AddAddress(authID int, address string, flagDefault bool) (*model.Address, error) {

	var newID int
	err := repo.DB.QueryRow(`
		INSERT INTO address (auth_id, address, flag_default)
		VALUES ($1, $2, $3)
		RETURNING id
	`, authID, address, flagDefault).Scan(&newID)

	if err != nil {
		repo.Logger.Error("Failed to add address", zap.Int("authID", authID), zap.String("address", address), zap.Error(err))
		return nil, fmt.Errorf("failed to add address: %v", err)
	}

	return &model.Address{
		ID:          newID,
		AuthID:      authID,
		Address:     address,
		FlagDefault: flagDefault,
	}, nil
}
