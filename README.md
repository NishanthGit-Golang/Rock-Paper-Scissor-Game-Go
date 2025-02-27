
# Rock-Paper-Scissors WebSocket Game Server

## Overview
This is a real-time multiplayer Rock-Paper-Scissors game server built in Go using WebSockets. The server handles player connections, matchmaking, game logic, and real-time updates.

## Features
- Player connects via WebSocket and is assigned a unique ID.
- Automatic matchmaking (pairs two players together).
- Players choose Rock, Paper, or Scissors.
- The server determines the winner and notifies both players.

### Installation & Running

1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Run the server:
   ```sh
   go run cmd/server/main.go
   ```
   The server will start at `ws://localhost:8080/ws`

---

##  Testing with Postman
### 1️. Connect to WebSocket
- Open Postman → **New Request** → **WebSocket Request**
- URL: `ws://localhost:8080/ws`
- Click **Connect**

### 2️. Register Player
Send JSON:
```json
{
  "name": "John"
}
```
Response:
```json
{
  "event": "connected",
  "playerId": "some-uuid"
}
```

### 3. Start a Match (Two Players Needed)
- Repeat step in a new WebSocket tab.
- Both players will receive:
  ```json
  {
    "event": "start_game",
    "opponentId": "opponent-uuid",
    "opponentName": "Alice"
  }
  ```

### 4. Send Choice (Rock, Paper, Scissors)
```json
{
  "event": "player_choice",
  "choice": "rock"
}
```

###  5. Get Game Result
Once both players submit their choice:
```json
{
  "event": "game_result",
  "winner": "John",
  "playerChoice": "rock",
  "opponentChoice": "scissors",
  "playerName": "John",
  "opponentName": "Alice"
}
