// internal/handlers/user_handler.go

package handlers

import (
    "blog/internal/models"
    "context"
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

var db *sql.DB

// SetDB sets the database connection to be used by handlers
func SetDB(database *sql.DB) {
    db = database
}

// GetUserHandler handles GET requests for retrieving a user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    if db == nil {
        log.Println("Database connection is not set")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    vars := mux.Vars(r)
    userID := vars["userID"]

    queries := models.New(db)
    user, err := queries.GetUser(context.Background(), userID)
    if err != nil {
        log.Println("Error retrieving user:", err)
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
