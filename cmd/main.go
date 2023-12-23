package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	dbAccessor "github.com/Coop25/the-meme-index/internal/accessors/database"
	dbManager "github.com/Coop25/the-meme-index/internal/managers/database"

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

	// Execute the CREATE TABLE IF NOT EXISTS statement
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS files (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			content BYTEA,
			tags TEXT[],
			url VARCHAR(255),
			description TEXT
		);
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fileAccessor := dbAccessor.New(db)
	fileManager := dbManager.New(fileAccessor)

	// Setup routes
	r := mux.NewRouter()

	r.HandleFunc("/upload", uploadHandler(fileManager)).Methods("POST")
	r.HandleFunc("/files/{id:[0-9]+}", getFileHandler(fileManager)).Methods("GET")
	// r.HandleFunc("/random", getRandomFileHandler(fileManager)).Methods("GET")
	// r.HandleFunc("/search", searchFilesHandler(fileManager)).Methods("GET")

	// Start the web server
	port := "8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
