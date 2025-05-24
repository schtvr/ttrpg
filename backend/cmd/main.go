package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	h "github.com/schtvr/ttrpg/backend/internal/handlers"
)

const port = ":1337"

func main() {

	// webscoket
	// hub := c.NewHub()
	// go hub.Run()
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	c.ServeWs(hub, w, r)
	// })

	r := mux.NewRouter()
	http.Handle("/", r)

	ch := h.NewCharacterHandler()
	r.HandleFunc("/character/generate", ch.GenerateNewCharacter).Methods("GET")
	r.HandleFunc("/character/new", ch.AddNewCharacter).Methods("POST")
	r.HandleFunc("/character/{id:[a-zA-Z0-9_-]+}", ch.GetCharacter).Methods("GET")

	r.HandleFunc("/characters", ch.GetCharacters).Methods("GET")
	r.HandleFunc("/characters/{id:[a-zA-Z0-9_-]+}", ch.UpdateCharacter).Methods("PUT")

	http.ListenAndServe(port, r)

	// log.Printf("websocket server running at port %s\n", port)
	log.Printf("server running at port %s\n", port)
}
