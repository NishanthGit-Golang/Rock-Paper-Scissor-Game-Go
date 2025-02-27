package game

import (
	"log"
	"sync"
)

type Player struct {
	ID     string
	Name   string
	Choice string
}

type Game struct {
	Player1 *Player
	Player2 *Player
	Mutex   sync.Mutex
}

func DetermineWinner(p1, p2 *Player) string {
	log.Printf("Determining winner between %s (%s) and %s (%s)", p1.ID, p1.Choice, p2.ID, p2.Choice)
	if p1.Choice == p2.Choice {
		log.Println("Result: Draw")
		return "draw"
	}
	rules := map[string]string{"rock": "scissors", "scissors": "paper", "paper": "rock"}
	if rules[p1.Choice] == p2.Choice {
		log.Printf("Winner: %s", p1.Name)
		return p1.Name
	}
	log.Printf("Winner: %s", p2.Name)
	return p2.Name
}
