package database

import (
	"github.com/Coop25/memeindex/internal/accessor/database"
)

type DatabaseManager struct {
	fileAccessor accessor.FileAccessor
}

func NewFileManager(fileAccessor accessor.FileAccessor) *DatabaseManager {
	return &DatabaseManager{fileAccessor: fileAccessor}
}
