package domain

import (
	"log"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Person struct {
	ID   string
	Conn *websocket.Conn
}

type Potato struct {
	HolderID  string
	PassCount int
	Timer     *time.Time
}

type Lobby struct {
	Players   []Person
	PlayerCh  chan Person   // A channel for each player
	StartGame chan struct{} //A signal to start a game
	Mutext    sync.Mutex
}

func NewLobby() *Lobby {
	return &Lobby{
		Players:   make([]Person, 0),
		PlayerCh:  make(chan Person, 5),
		StartGame: make(chan struct{}),
	}
}

func (lb *Lobby) MaxCount() bool {
	return len(lb.Players) == 3
}
func (lb *Lobby) AddToLobby(p Person) {
	lb.Players = append(lb.Players, p)
}

func (lb *Lobby) Broadcast(message string) {
	lb.Mutext.Lock()
	defer lb.Mutext.Unlock()
	for _, player := range lb.Players {
		err := player.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Error sending to player %v: %v", player.ID, err)
		}
	}
}

func (lb *Lobby) EliminatePlayer(playerID string) {
	lb.Mutext.Lock()
	defer lb.Mutext.Unlock()
	for i, p := range lb.Players {
		if p.ID == playerID {
			lb.Players = append(lb.Players[:i], lb.Players[i+1:]...)
			p.Conn.Close()
			break
		}
	}
}

func (lb *Lobby) StartGameLoop() {
	log.Printf("Starting game loop")
	potato := Potato{}

	go func() {
		for {
			select {
			case <-lb.StartGame:
				//Assign the potato to one player
				it := rand.IntN(len(lb.Players))
				potato.HolderID = string(lb.Players[it].ID)
				log.Printf("The player holding the potato is: %v", potato.HolderID)
				timer := time.After(30 * time.Second)
				mvTimer := time.NewTicker(5 * time.Second)
				go func() {
					for {
						if len(lb.Players) == 1 {
							lb.Broadcast("Player " + lb.Players[0].ID + " is the winner!")
							return
						}
						it = rand.IntN(len(lb.Players))
						select {
						case <-timer:
							log.Printf("Loser player: %v", potato.HolderID)
							lb.Broadcast("Player " + potato.HolderID + " is eliminated!")
							lb.EliminatePlayer(potato.HolderID)
							return
						case <-mvTimer.C:
							potato.HolderID = string(lb.Players[it].ID)
							lb.Broadcast("Player " + potato.HolderID + " is holding the potato")
							log.Printf("New player holding the potato is: %v", potato.HolderID)

						}
					}
				}()

			}
		}
	}()

}
