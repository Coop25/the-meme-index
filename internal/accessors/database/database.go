package database

import (
	"database/sql"
)

type databaseAccessor struct {
	db *sql.DB
}

func New(db *sql.DB) DatabaseAccessor {
	return &databaseAccessor{db: db}
}
