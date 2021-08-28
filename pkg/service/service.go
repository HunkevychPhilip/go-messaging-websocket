package service

import "github.com/gorilla/websocket"

type Chat interface {
	NewClient(conn *websocket.Conn)
	Router()
}

type Service struct {
	ChatService Chat
}

func NewService() *Service {
	return &Service{
		ChatService: NewChatService(),
	}
}
