// internal/routes/routes.go

package routes

import (
    "net/http"
)

// Init initializes HTTP routes
func Init() {
    // Define routes here
    http.HandleFunc("/", handler)
}

// Handler is your HTTP request handler
func handler(w http.ResponseWriter, r *http.Request) {
    // Respond with "Hello, World!"
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}
