// package main is the entry point for the web application.
package main

import (
	"github.com/calmdev/go-docker-from-scratch/internal/config"
	"github.com/calmdev/go-docker-from-scratch/internal/server"
)

func main() {
	// Create a new configuration.
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Create a new server.
	if err := server.NewServer(config); err != nil {
		panic(err)
	}
}
