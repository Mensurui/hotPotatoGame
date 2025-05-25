package main

import (
	"fmt"
	"log"
	"net/http"

	sessionmanager "github.com/Mensurui/hotPotatoGame/internal/sessionManager"
	"github.com/Mensurui/hotPotatoGame/internal/transport/ws/handler"
)

func main() {
	sess := sessionmanager.NewSessionManager()
	sess.ListenToLobby()
	sess.Lobby.StartGameLoop()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsh := handler.NewWebSocketHandler(*sess)
		wsh.WSHandler(w, r)
	})
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error establishing connection: %v\n", err)
	}
	fmt.Println("Server is listening...")
}
