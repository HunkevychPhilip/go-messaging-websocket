package service

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/datastore"
	"github.com/gorilla/websocket"
	"net/http"
)

type Chat interface {
	UpgradeConn(http.ResponseWriter, *http.Request) (*websocket.Conn, error)
	ServeNewConn(*websocket.Conn)
}

type Service struct {
	Chat
}

func NewService(db *datastore.Datastore) *Service {
	return &Service{
		Chat: NewChatService(db.Chat),
	}
}
