package clientmanager

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients = make(map[string]*websocket.Conn)
	mutex   sync.Mutex
)

func RegisterClient(playerID string, conn *websocket.Conn) {
	mutex.Lock()
	clients[playerID] = conn
	mutex.Unlock()
}

func GetClient(playerID string) *websocket.Conn {
	mutex.Lock()
	defer mutex.Unlock()
	return clients[playerID]
}
