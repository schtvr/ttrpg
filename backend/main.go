package main

import (
	"log"
	"net/http"

	"github.com/schtvr/ttrpg/backend/middleware"
	"github.com/schtvr/ttrpg/backend/routes"
)

func main() {
	// Initialize the router
	mux := http.NewServeMux()

	// Load routes
	routes.RegisterRoutes(mux)

	// Wrap with middleware
	handler := middleware.ApplyMiddleware(mux)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
