package database

import (
	dbAccessor "github.com/Coop25/the-meme-index/internal/accessors/database"
)

func (m *databaseManager) UploadFile(in dbAccessor.NewMeme) (int, error) {
	return m.fileAccessor.UploadFile(in)
}

func (m *databaseManager) GetFileByID(id int) (dbAccessor.Meme, error) {
	return m.fileAccessor.GetFileByID(id)
}

func (m *databaseManager) GetRandomFile() (int, string, error) {
	panic("not Implemented")
}

func (m *databaseManager) SearchFilesByTags(tagList []string) ([]int, []string, error) {
	panic("not Implemented")
}
