// internal/routes/routes.go

package routes

import (
	"blog/internal/handlers"
	"blog/internal/handlers/auth"

	"github.com/gorilla/mux"
)

// Init initializes HTTP routes
func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/special", handlers.HelloHandler).Methods("GET")
	router.HandleFunc("/users/{userID}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/register", auth.RegisterHandler).Methods("POST")
	return router
}
