package handler

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) Chat(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err.Error())

		return
	}
	log.Println("Client successfully connected.")

	h.service.ChatService.NewClient(conn)
}
