// package server is responsible for starting the HTTP server and handling requests.
package server

import (
	"net/http"

	"github.com/calmdev/go-docker-from-scratch/internal/config"
)

// NewServer starts a new server.
func NewServer(config *config.Config) error {
	// Create a new HTTP server.
	mux := http.NewServeMux()

	// Register the handler function for the index route.
	mux.HandleFunc("/", hello)

	// Start the HTTP server.
	return http.ListenAndServe(config.HttpListenAddress, mux)
}
