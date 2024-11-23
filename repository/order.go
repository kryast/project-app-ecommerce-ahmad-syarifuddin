package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (repo *ProductRepositoryDB) GetOrderData(authID int) (*model.Order, error) {
	carts, err := repo.GetAllCart()
	if err != nil {
		repo.Logger.Error("Error fetching cart items", zap.Error(err), zap.Int("auth_id", authID))
		return nil, fmt.Errorf("failed to fetch cart items: %w", err)
	}

	var address string
	queryAddress := `SELECT address FROM address WHERE auth_id = $1`
	err = repo.DB.QueryRow(queryAddress, authID).Scan(&address)
	if err != nil {
		repo.Logger.Error("Error fetching address", zap.Error(err), zap.Int("auth_id", authID))
		return nil, fmt.Errorf("failed to fetch address: %w", err)
	}

	order := &model.Order{
		AuthID:  authID,
		Address: address,
		Payment: "Bank Jaga ",
		Cart:    carts,
	}

	return order, nil
}
