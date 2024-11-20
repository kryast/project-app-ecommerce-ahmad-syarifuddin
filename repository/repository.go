package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	AuthRepo    AuthRepositoryDB
	ProductRepo ProductRepositoryDB
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		AuthRepo:    NewAuthRepository(db, log),
		ProductRepo: NewProductRepository(db, log),
	}
}
