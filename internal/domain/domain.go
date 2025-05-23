package domain

import (
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
