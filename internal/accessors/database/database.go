package database

import (
	"database/sql"
	"io"
)

type DatabaseAccessor interface {
	UploadFile(name string, content io.Reader, tags, url string) (int, error)
	GetFileByID(fileID int) ([]byte, string, string, error)
	GetRandomFile() (int, string, error)
	SearchFilesByTags(tagList []string) ([]int, []string, error)
}

type DatabaseAccessor struct {
	db *sql.DB
}

func NewDatabaseAccessor(db *sql.DB) DatabaseAccessor {
	return &DatabaseAccessor{db: db}
}