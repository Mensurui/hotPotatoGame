package handler

import (
	"fmt"
	"net/http"

	"github.com/Mensurui/hotPotatoGame/internal/domain"
	"github.com/Mensurui/hotPotatoGame/internal/helper"
	sessionmanager "github.com/Mensurui/hotPotatoGame/internal/sessionManager"
)

type WebSocketHandler struct {
	SessionMgr sessionmanager.SessionManager
}

func NewWebSocketHandler(sess sessionmanager.SessionManager) *WebSocketHandler {
	return &WebSocketHandler{
		SessionMgr: sess,
	}
}

func (wsh *WebSocketHandler) WSHandler(w http.ResponseWriter, r *http.Request) {
	//Whenever this handler is called we add the user to the sm
	conn, err := helper.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	person := domain.Person{
		ID:   r.URL.Query().Get("id"),
		Conn: conn,
	}
	wsh.SessionMgr.AddPlayerToSession(&person)
	fmt.Println("New connection established with ID:", person.ID)
	fmt.Println("Connection:", person.Conn.LocalAddr().String())

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received message: %s\n", msg)
	}
}
