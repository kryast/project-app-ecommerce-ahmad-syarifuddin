package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

func (r *AuthRepositoryDB) Login(login *model.Login) (*model.Register, error) {
	var user model.Register
	var storedPassword string

	if login.Email != "" {
		query := `SELECT id, name, email, phone, password FROM auths WHERE email = $1`
		err := r.DB.QueryRow(query, login.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("email not found")
			}
			r.Logger.Error("Error querying email", zap.Error(err))
			return nil, fmt.Errorf("error querying email: %w", err)
		}
	} else if login.Phone != "" {
		query := `SELECT id, name, email, phone, password FROM auths WHERE phone = $1`
		err := r.DB.QueryRow(query, login.Phone).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("phone number not found")
			}
			r.Logger.Error("Error querying phone", zap.Error(err))
			return nil, fmt.Errorf("error querying phone: %w", err)
		}
	} else {
		return nil, fmt.Errorf("email or phone must be provided")
	}

	if login.Password != storedPassword {
		return nil, fmt.Errorf("incorrect password")
	}

	return &user, nil
}
