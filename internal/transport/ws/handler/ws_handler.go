package handler

import (
	"fmt"
	"net/http"

	"github.com/Mensurui/hotPotatoGame/internal/domain"
	"github.com/Mensurui/hotPotatoGame/internal/helper"
)

func WSHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("New connection established with ID:", person.ID)
	fmt.Println("Connection:", person.Conn.LocalAddr().String())
}
