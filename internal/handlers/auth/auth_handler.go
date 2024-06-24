// internal/handlers/auth/auth_handler.go

package auth

import (
	"blog/internal/models"
	"blog/internal/utils"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var db *sql.DB
var queries *models.Queries
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// SetDB sets the database connection for the handlers
func SetDB(database *sql.DB) {
	db = database
	queries = models.New(db)
}

// RegisterHandler handles the user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username   string `json:"username" validate:"required,min=3,max=32"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required,min=6"`
		RePassword string `json:"re_password" validate:"required,eqfield=Password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Check if the email already exists in the database
	if utils.EmailExists(request.Email) {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Create the new user
	userParams := models.CreateUserParams{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		FullName: sql.NullString{String: "", Valid: false},
		DateOfBirth: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}

	ctx := context.Background()
	if err := queries.CreateUser(ctx, userParams); err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
