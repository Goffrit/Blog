// internal/utils/email.go

package utils

import (
	"blog/internal/models"
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

// SetDB sets the database connection to be used by handlers
func SetDB(database *sql.DB) {
	db = database
}

// EmailExists checks if the given email already exists in the database
func EmailExists(email string) bool {
	// Check if db is nil
	if db == nil {
		log.Println("Database connection is not initialized")
		return true // Assume email exists to avoid registration or handle as necessary
	}

	queries := models.New(db)
	_, err := queries.GetUserByEmail(context.Background(), email)
	if err == sql.ErrNoRows {
		return false // Email does not exist
	} else if err != nil {
		log.Printf("Error checking email existence: %v", err)
		return true // Assume email exists to avoid registration or handle as necessary
	}
	return true // Email exists
}
