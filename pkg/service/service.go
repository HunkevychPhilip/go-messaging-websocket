package service

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/messaging"
)

type Service struct {
	Chat
}

func NewService(ps *messaging.Messaging) *Service {
	return &Service{
		Chat: NewChatService(ps),
	}
}
