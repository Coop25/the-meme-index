package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dbManager "github.com/Coop25/the-meme-index/internal/managers/database"
	"github.com/gorilla/mux"
)

func getFileHandler(fileManager dbManager.DatabaseManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the file ID from the request parameters
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid file ID", http.StatusBadRequest)
			return
		}

		// Run Manager Layer
		meme, err := fileManager.GetFileByID(id)
		if err != nil {
			http.Error(w, "Error handling database operations", http.StatusInternalServerError)
			return
		}

		// Set the appropriate headers for file download
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", meme.Name))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.Itoa(len(meme.Content)))

		// Write the file content to the response
		if _, err := w.Write(meme.Content); err != nil {
			http.Error(w, "Error writing file to response", http.StatusInternalServerError)
			return
		}

		// Prepare JSON data for filename, tags, and URL
		fileInfo := struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Tags        []string `json:"tags"`
			URL         string `json:"url"`
			Description string `json:"description"`
		}{
			Id:          meme.Id,
			Name:        meme.Name,
			Tags:        meme.Tags,
			URL:         meme.Url,
			Description: meme.Id,
		}

		// Marshal the struct to JSON
		fileInfoJSON, err := json.Marshal(fileInfo)
		if err != nil {
			http.Error(w, "Error encoding file info to JSON", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header for JSON
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON data to the response
		if _, err := w.Write(fileInfoJSON); err != nil {
			http.Error(w, "Error writing JSON data to response", http.StatusInternalServerError)
			return
		}
	}
}
