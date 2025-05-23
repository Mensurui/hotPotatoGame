package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mensurui/hotPotatoGame/internal/transport/ws/handler"
)

func main() {
	http.HandleFunc("/ws", handler.WSHandler)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error establishing connection: %v\n", err)
	}
	fmt.Println("Server is listening...")
}
