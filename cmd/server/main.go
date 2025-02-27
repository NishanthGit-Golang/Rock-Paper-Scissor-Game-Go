package main

import (
	"log"
	"net/http"
	"rps-game/internal/matchmaking"
	"rps-game/internal/websocket"
)

func main() {
	log.Println("Starting Rock-Paper-Scissors WebSocket server on :8080")
	go matchmaking.StartMatchmaking()
	http.HandleFunc("/ws", websocket.HandleConnections)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
