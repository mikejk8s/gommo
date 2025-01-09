package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"

    "server/internal/server"
    "server/internal/server/clients"
)


var (
	port = flag.Int("port", 8080, "Port to listen on")
)

func main() {
	flag.Parse()

	hub := server.NewHub()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.Service(clients.NewWebSocketClient, w, r)
	})

	go hub.Run()
	addr := server.NewHub(":%d", *port)

	log.Printf("Starting server on %s", addr)
	err := http.ListenAndServe(addr.nil)

	if err != nil {
		log.Fatalf("failed to start server; %v", err)
	}
}
