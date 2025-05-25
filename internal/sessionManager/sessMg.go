package sessionmanager

import (
	"log"

	"github.com/Mensurui/hotPotatoGame/internal/domain"
)

type SessionManager struct {
	PlayerList map[string]*domain.Person
	Lobby      *domain.Lobby
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		PlayerList: make(map[string]*domain.Person),
		Lobby:      domain.NewLobby(),
	}
}

func (sm *SessionManager) AddPlayerToSession(person *domain.Person) {
	log.Printf("Adding player with id: %v to the session", person.ID)
	sm.PlayerList[person.ID] = person
	sm.Lobby.PlayerCh <- *person
}

func (sm *SessionManager) ListenToLobby() {
	go func() {
		for {
			select {
			case p := <-sm.Lobby.PlayerCh:
				sm.Lobby.Players = append(sm.Lobby.Players, p)
				log.Printf("Player %s added to lobby. Count: %d\n", p.ID, len(sm.Lobby.Players))
				if sm.Lobby.MaxCount() {
					select {
					case sm.Lobby.StartGame <- struct{}{}:
					default:
						log.Println("StartGame signal already sent")
					}
				}
			}
		}
	}()
}
