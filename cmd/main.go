package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"memeindex/internal/accessor"
	manager "memeindex/internal/managers/database"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "user=username dbname=mydatabase sslmode=disable" // replace with your PostgreSQL connection string
)

func main() {
	// Connect to the PostgreSQL database
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	fileAccessor := accessor.NewFileAccessor(db)
	fileManager := manager.NewFileManager(fileAccessor)

	// Setup routes
	r := mux.NewRouter()

	r.HandleFunc("/upload", uploadHandler(fileManager)).Methods("POST")
	r.HandleFunc("/files/{id:[0-9]+}", getFileHandler(fileManager)).Methods("GET")
	r.HandleFunc("/random", getRandomFileHandler(fileManager)).Methods("GET")
	r.HandleFunc("/search", searchFilesHandler(fileManager)).Methods("GET")

	// Start the web server
	port := "8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
