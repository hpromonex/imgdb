package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database
	db, err := NewDB()
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Initialize the handler with the database connection
	handler := &Handler{DB: db}

	// Set up routes
	mux := http.NewServeMux()
	mux.Handle("/images/upload", http.HandlerFunc(handler.UploadImage))
	mux.Handle("/images/search", http.HandlerFunc(handler.SearchImagesByTags))
	mux.Handle("/images/edit-metadata", http.HandlerFunc(handler.EditImageMetadata))
	mux.Handle("/images/serve", http.HandlerFunc(handler.ServeImages))

	staticFileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", staticFileServer)

	// Start the server
	log.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
