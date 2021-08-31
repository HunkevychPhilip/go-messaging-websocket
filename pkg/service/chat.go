package service

import (
	"encoding/json"
	"fmt"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/messaging"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/websocket"
	"strings"
)

type Chat interface {
	NewClient(conn *websocket.Conn)
}

type ChatService struct {
	messaging *messaging.Messaging
}

func NewChatService(m *messaging.Messaging) *ChatService {
	return &ChatService{
		messaging: m,
	}
}

func (c *ChatService) NewClient(conn *websocket.Conn) {
	nickname := gofakeit.Name()

	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection" + err.Error())
		}
	}()

	p := c.messaging.Client.PSubscribe("messages.*")

	_, err := p.Receive()
	if err != nil {
		fmt.Println("Failed to subscribe" + err.Error())

		return
	}

	ch := p.Channel()

	go func() {
		for msg := range ch {
			fmt.Printf("Received message: %s.\n", msg.Channel)

			switch msg.Channel {
			case "messages.event.new":
				var message types.Message

				if err := json.NewDecoder(strings.NewReader(msg.Payload)).Decode(&message); err != nil {
					continue
				}

				err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("> %s: %s", message.Owner, message.Content)))
				if err != nil {
					fmt.Println("Failed to write message" + err.Error())

					return
				}
			}
		}
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message" + err.Error())

			return
		}

		message := types.Message{
			Owner:   nickname,
			Content: string(p),
		}

		p, err = json.Marshal(message)
		if err != nil {
			fmt.Println("Failed to marshal message.")

			return
		}

		res := c.messaging.Client.Publish("messages.event.new", p)
		if res.Err() != nil {
			fmt.Println("Failed to publish message")

			return
		}
	}
}
