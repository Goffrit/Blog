// internal/routes/routes.go

package routes

import (
    "blog/internal/handlers"
    "github.com/gorilla/mux"
)

// Init initializes HTTP routes
func Init() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/users/{userID}", handlers.GetUserHandler).Methods("GET")
    router.HandleFunc("/users/special", handlers.HelloHandler).Methods("GET")
    return router
}