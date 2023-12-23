package database

import (
	dbAccessor "github.com/Coop25/the-meme-index/internal/accessors/database"
)

type DatabaseManager interface {
	UploadFile(dbAccessor.NewMeme) (int, error)
	GetFileByID(id int) (dbAccessor.Meme, error)
	GetRandomFile() (int, string, error)
	SearchFilesByTags(tagList []string) ([]int, []string, error)
}
