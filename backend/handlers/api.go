package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, TTRPG backend!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
