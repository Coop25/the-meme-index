package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	dbAccessor "github.com/Coop25/the-meme-index/internal/accessors/database"
	dbManager "github.com/Coop25/the-meme-index/internal/managers/database"
)

func uploadHandler(fileManager dbManager.DatabaseManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form data
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the file from the form data
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to get file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Read the content of the uploaded file
		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read file content", http.StatusInternalServerError)
			return
		}

		// Extract tags, URL, and description from the form data (adjust field names accordingly)
		newMeme := dbAccessor.NewMeme{
			Name:        handler.Filename,
			Content:     content,
			Tags:        strings.Split(r.FormValue("tags"), ","),
			Url:         r.FormValue("url"),
			Description: r.FormValue("description"),
		}

		// Call the file manager to upload file information to the database
		fileID, err := fileManager.UploadFile(newMeme)
		if err != nil {
			http.Error(w, "Error handling database operations", http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		response := map[string]interface{}{
			"fileID":  fileID,
			"message": "File uploaded successfully, and information stored in the database",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
