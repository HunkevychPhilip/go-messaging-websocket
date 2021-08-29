package service

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/gorilla/websocket"
)

type Chat interface {
	NewClient(conn *websocket.Conn)
}

type Service struct {
	ChatService Chat
}

func NewService(c *types.ChatChannels) *Service {
	return &Service{
		ChatService: NewChatService(c),
	}
}
