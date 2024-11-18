package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	SampelRepo SampelRepository
	AuthRepo   AuthRepositoryDB
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		SampelRepo: NewSampelRepository(db, log),
		AuthRepo:   NewAuthRepository(db, log),
	}
}
