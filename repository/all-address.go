package repository

import (
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"
)

func (r *AuthRepositoryDB) GetAddressesByAuthID(authID int) ([]model.Address, error) {
	query := `
		SELECT auth_id, address, flag_default
		FROM address
		WHERE auth_id = $1
	`

	rows, err := r.DB.Query(query, authID)
	if err != nil {
		return nil, fmt.Errorf("failed to query addresses: %v", err)
	}
	defer rows.Close()

	var addresses []model.Address

	for rows.Next() {
		var address model.Address
		if err := rows.Scan(&address.AuthID, &address.Address, &address.FlagDefault); err != nil {
			return nil, fmt.Errorf("failed to scan address: %v", err)
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
