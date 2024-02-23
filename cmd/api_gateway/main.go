package main

import (
	"log"

	"github.com/sgokul961/echo-hub-api-gateway/pkg/config"
	"github.com/sgokul961/echo-hub-api-gateway/pkg/server"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize server
	srv := server.NewServer(&cfg)

	// Start server
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
