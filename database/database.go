package database

import (
	"github.com/bektosh/fiber-app/database/postgres"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Psql *postgres.Postgres
}

func New(db *sqlx.DB) *Storage {
	return &Storage{Psql: postgres.New(db)}
}
