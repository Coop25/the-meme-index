package database

import (
	dbAccessor "github.com/Coop25/the-meme-index/internal/accessors/database"
)

type databaseManager struct {
	fileAccessor dbAccessor.DatabaseAccessor
}

func New(fileAccessor dbAccessor.DatabaseAccessor) DatabaseManager {
	return &databaseManager{fileAccessor: fileAccessor}
}
