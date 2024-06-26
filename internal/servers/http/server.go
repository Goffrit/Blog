// internal/servers/http/server.go

package http

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"blog/internal/handlers"
	"blog/internal/handlers/auth"
	"blog/internal/routes"
	"blog/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// StartServer initializes and starts the HTTP server
func StartServer() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Get the driver from environment variables
	driver := os.Getenv("DRIVER")
	if driver == "" {
		log.Fatal("DRIVER environment variable is not set")
	}

	// Get the data source from environment variables
	dataSource := os.Getenv("DATA_SOURCE")
	if dataSource == "" {
		log.Fatal("DATA_SOURCE environment variable is not set")
	}

	// Open SQLite database
	db, err = sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set the database connection for handlers
	handlers.SetDB(db)

	// Set the database connection for handlers
	auth.SetDB(db)

	//Set the database connection for utils package
	utils.SetDB(db)

	// Initialize routes
	router := routes.Init()

	// Start HTTP server
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
