package handlers

import (
	"net/http"
)

// HelloHandler is your HTTP request handler for "/"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with "Hello, World!"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
