package service

import (
	"fmt"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/websocket"
)

type ChatService struct {
	channels *types.ChatChannels
}

func NewChatService(c *types.ChatChannels) *ChatService {
	return &ChatService{
		channels: c,
	}
}

func (c *ChatService) NewClient(conn *websocket.Conn) {
	msgChan := make(chan *types.Msg, 100)

	randNick := gofakeit.Name()

	c.channels.ClientRequests <- &types.Client{
		Nickname: randNick,
		MsgChan:    msgChan,
	}
	defer func() { c.channels.ClientDisconnects <- randNick }()

	go func() {
		for msg := range msgChan {
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Text))
			if err != nil {
				fmt.Println("Failed to write message:", err.Error())

				return
			}
		}
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error:", err.Error())

			return
		}

		c.channels.Messages <- &types.Msg{
			ClientNick: randNick,
			Text:       string(p),
		}
	}
}
