package database

func (a *databaseAccessor) GetFileByID(fileID int) (Meme, error) {
	// Retrieve the file content from the database
	query := `SELECT id, name, content, tags, url, description FROM files WHERE id = $1;`

	var meme Meme
	err := a.db.QueryRow(query, fileID).Scan(meme.Name, meme.Content, meme.Tags, meme.Url, meme.Description)
	if err != nil {
		return Meme{}, err
	}

	return meme, nil
}
