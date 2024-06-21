// internal/routes/routes.go

package routes

import (
	"net/http"

	"blog/internal/handlers" // Import handlers package
)

// Init initializes HTTP routes
func Init() {
	// Define routes here
	http.HandleFunc("/", handlers.HelloHandler)
}
