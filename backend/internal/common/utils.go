package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func EncodeAndSend(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	if err := e.Encode(data); err != nil {
		log.Printf("error encoding json: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
