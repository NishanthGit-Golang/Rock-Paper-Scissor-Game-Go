package websocket

import (
	"log"
	"net/http"
	"rps-game/internal/clientmanager"
	"rps-game/internal/game"
	"rps-game/internal/matchmaking"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("New WebSocket connection request received")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	// Read the first message to get the player name
	var msg map[string]string
	err = conn.ReadJSON(&msg)
	if err != nil {
		log.Println("Error reading player name:", err)
		conn.Close()
		return
	}

	playerName, ok := msg["name"]
	if !ok || playerName == "" {
		playerName = "Player-" + uuid.New().String()[:5] // Default name if not provided
	}

	playerID := uuid.New().String()
	log.Printf("Player connected: %s (ID: %s)", playerName, playerID)

	player := &game.Player{ID: playerID, Name: playerName}
	clientmanager.RegisterClient(playerID, conn)

	matchmaking.AddToQueue(player)
	listenForMessages(conn, player)
}

func listenForMessages(conn *websocket.Conn, player *game.Player) {
	for {
		var msg map[string]string
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message from player %s: %v", player.ID, err)
			break
		}

		log.Printf("Received message from player %s: %v", player.ID, msg)
		if choice, ok := msg["choice"]; ok {
			player.Choice = choice
			log.Printf("Player %s chose: %s", player.ID, choice)
			g, exists := matchmaking.GetGame(player.ID)
			if exists {
				go matchmaking.ProcessGame(g)
			}
		}
	}
}
