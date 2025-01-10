package routes

import (
	"net/http"

	"github.com/schtvr/ttrpg/backend/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	// API routes
	mux.HandleFunc("/api/hello", handlers.HelloWorldHandler)

	// WebSocket routes
	mux.HandleFunc("/ws", handlers.WebSocketHandler)
}
