package main

import (
	"flag"
	"log"
	"net/http"

	c "github.com/schtvr/ttrpg/backend/internal/clients"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	hub := c.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c.ServeWs(hub, w, r)
	})
	log.Println("starting websocket server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("websocket server running at port 8080")
}
