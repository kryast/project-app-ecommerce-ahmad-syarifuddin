package repository

import (
	"fmt"
)

func (r *AuthRepositoryDB) StoreSessionToken(adminID int, token string) error {
	query := `UPDATE auths SET token = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, token, adminID)
	if err != nil {
		return fmt.Errorf("failed to store session token: %v", err)
	}
	return nil
}
