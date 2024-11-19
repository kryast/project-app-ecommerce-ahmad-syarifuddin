package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	SampelRepo  SampelRepository
	AuthRepo    AuthRepositoryDB
	ProductRepo ProductRepositoryDB
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		SampelRepo:  NewSampelRepository(db, log),
		AuthRepo:    NewAuthRepository(db, log),
		ProductRepo: NewProductRepository(db, log),
	}
}
