package matchmaking

import (
	"log"
	"rps-game/internal/clientmanager"
	"rps-game/internal/game"
	"sync"
	"time"
)

var (
	waitingQueue = make(chan *game.Player, 10)
	Games        = make(map[string]*game.Game)
	mutex        sync.Mutex
)

func AddToQueue(player *game.Player) {
	log.Printf("Player %s added to matchmaking queue", player.ID)
	waitingQueue <- player
}

func StartMatchmaking() {
	for {
		p1 := <-waitingQueue
		p2 := <-waitingQueue
		log.Printf("Match found: %s vs %s", p1.ID, p2.ID)

		newGame := &game.Game{Player1: p1, Player2: p2}
		mutex.Lock()
		Games[p1.ID] = newGame
		Games[p2.ID] = newGame
		mutex.Unlock()

		go ProcessGame(newGame)
	}
}

func ProcessGame(g *game.Game) {
	for {
		g.Mutex.Lock()
		if g.Player1.Choice != "" && g.Player2.Choice != "" {
			winner := game.DetermineWinner(g.Player1, g.Player2)
			log.Printf("Game result: Winner - %s", winner)
			notifyPlayers(g, winner)
			g.Mutex.Unlock()
			return
		}
		g.Mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}

func notifyPlayers(g *game.Game, winner string) {
	response := map[string]string{
		"event":          "game_result",
		"winner":         winner,
		"playerChoice":   g.Player1.Choice,
		"opponentChoice": g.Player2.Choice,
		"playerName":     g.Player1.Name,
		"opponentName":   g.Player2.Name,
	}

	for _, player := range []*game.Player{g.Player1, g.Player2} {
		conn := clientmanager.GetClient(player.ID)
		if conn != nil {
			conn.WriteJSON(response)
		}
	}
}

func GetGame(playerID string) (*game.Game, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	g, exists := Games[playerID]
	return g, exists
}
