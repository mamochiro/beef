package postgres

import "github.com/mamochiro/beef/internals/infrastructure/database"

type PostgresRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Interface {
	return &PostgresRepository{db: db}
}
