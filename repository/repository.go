package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	SampelRepo SampelRepository
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		SampelRepo: NewSampelRepository(db, log),
	}
}
