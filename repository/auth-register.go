package repository

import (
	"database/sql"
	"fmt"
	"project-app-ecommerce-ahmad-syarifuddin/model"

	"go.uber.org/zap"
)

type AuthRepositoryDB struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewAuthRepository(db *sql.DB, Log *zap.Logger) AuthRepositoryDB {
	return AuthRepositoryDB{
		DB:     db,
		Logger: Log,
	}
}

func (r *AuthRepositoryDB) Register(register *model.Register) error {
	query := `
		INSERT INTO auths (name, email, phone, password)
		VALUES ($1, 
				COALESCE(NULLIF($2, ''), NULL), 
				COALESCE(NULLIF($3, ''), NULL), 
				$4);
		`

	_, err := r.DB.Exec(query, register.Name, register.Email, register.Phone, register.Password)
	if err != nil {
		r.Logger.Error("Failed to insert new register", zap.Error(err))
		return fmt.Errorf("failed to insert new register: %w", err)
	}

	return nil
}

func (r *AuthRepositoryDB) CheckEmailOrPhoneExists(register *model.Register) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*) 
		FROM auths
		WHERE email = $1 OR phone = $2;
	`
	err := r.DB.QueryRow(query, register.Email, register.Phone).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check email or phone: %w", err)
	}

	return count > 0, nil
}
