// internal/handlers/user_handler.go

package handlers

import (
 
 
 
    "regexp"
    "time"


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

    queries := models.New(db)
    ctx := context.Background()
    if err := queries.CreateUser(ctx, userParams); err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
