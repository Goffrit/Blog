package http

import (
    "database/sql"
    "log"
    "net/http"

	"blog/internal/routes" 

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// StartServer initializes and starts the HTTP server
func StartServer() {
    // Open SQLite database
    var err error
    db, err = sql.Open("sqlite3", "file:config/sqlite.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // Initialize routes from the routes package
    routes.Init()
	
    // Start HTTP server
    log.Println("Starting server on port 8181...")
    if err := http.ListenAndServe(":8181", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

