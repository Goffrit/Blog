// internal/handlers/auth/auth_handler.go

package auth

import (
	"blog/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"
)

var db *sql.DB
var queries *models.Queries

// SetDB sets the database connection for the handlers
func SetDB(database *sql.DB) {
	db = database
	queries = models.New(db)
}

// RegisterHandler handles the user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		RePassword string `json:"re_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if request.Password != request.RePassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	if match, _ := regexp.MatchString(emailRegex, request.Email); !match {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Create the new user
	userParams := models.CreateUserParams{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password, // You should hash the password before storing it
		FullName: sql.NullString{String: "", Valid: false},
		DateOfBirth: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}

	ctx := context.Background()
	if err := queries.CreateUser(ctx, userParams); err != nil {
		log.Printf("Failed to create user: %v", err) // Log the specific error
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
