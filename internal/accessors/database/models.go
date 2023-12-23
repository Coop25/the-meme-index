package database

type DatabaseAccessor interface {
	UploadFile(NewMeme) (int, error)
	GetFileByID(fileID int) (Meme, error)
	GetRandomFile() (int, string, error)
	SearchFilesByTags(tagList []string) ([]int, []string, error)
}

type NewMeme struct {
	Name        string
	Content     []byte
	Tags        []string
	Url         string
	Description string
}

type Meme struct {
	Id          string
	Name        string
	Content     []byte
	Tags        []string
	Url         string
	Description string
}
