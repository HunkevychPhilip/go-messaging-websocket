package service

import (
	"fmt"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/websocket"
	"log"
)

var (
	clientRequests    = make(chan *types.NewClient, 100)
	clientDisconnects = make(chan string, 100)
	messages          = make(chan *types.Msg, 100)
)

type ChatService struct {
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (c *ChatService) NewClient(conn *websocket.Conn) {
	msgChan := make(chan *types.Msg, 100)
	randNick := gofakeit.Name()
	clientRequests <- &types.NewClient{
		ClientNick: randNick,
		MsgChan:    msgChan,
	}
	defer func() { clientDisconnects <- randNick }()

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

		messages <- &types.Msg{
			ClientNick: randNick,
			Text:       string(p),
		}
	}
}

func (c *ChatService) Router() {
	clients := make(map[string]chan *types.Msg)

	for {
		select {
		case req := <-clientRequests:
			clients[req.ClientNick] = req.MsgChan
			log.Println("Websocket connected: " + req.ClientNick)
		case ClientNick := <-clientDisconnects:
			close(clients[ClientNick])
			delete(clients, ClientNick)
			log.Println("Websocket disconnected: " + ClientNick)
		case msg := <-messages:
			for _, msgChan := range clients {
				if len(msgChan) < cap(msgChan) {
					msgChan <- msg
				}
			}
		}
	}
}
