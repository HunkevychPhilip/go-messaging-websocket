package service

import (
	"encoding/json"
	"fmt"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/datastore"
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	messagesChannelPrefix = "messages.*"
	messagesEventNew      = "messages.event.new"

	errWebsocketGoingAway = "1001"
)

var upgradeTmpl = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ChatService struct {
	rdbChat datastore.Chat
}

func NewChatService(dc datastore.Chat) *ChatService {
	return &ChatService{
		rdbChat: dc,
	}
}

func (c *ChatService) UpgradeConn(rw http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgradeTmpl.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgradeTmpl.Upgrade(rw, r, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *ChatService) ServeNewConn(conn *websocket.Conn) {
	nickname := gofakeit.Name()

	pubsub, err := c.rdbChat.Subscribe(messagesChannelPrefix)
	if err != nil {
		logrus.Error(err)

		return
	}

	conn.SetCloseHandler(func(code int, text string) error {
		logrus.Infof("Connection closed for user %s", nickname)

		err = pubsub.Unsubscribe()
		if err != nil {
			return err
		}
		err = pubsub.Close()
		if err != nil {
			return err
		}

		return nil
	})

	go func() {
		for msg := range pubsub.Channel() {
			logrus.Infof("Received message: %s.", msg.Channel)

			switch msg.Channel {
			case "messages.event.new":
				if err = websocketWrite(conn, msg.Payload); err != nil {
					logrus.Error(err)

					return
				}
			default:
				logrus.Infof("Unrecognized event: %s.", msg.Channel)
			}
		}
	}()

	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			if strings.Contains(err.Error(), errWebsocketGoingAway) {
				logrus.Info("Connection gone away...")

				return
			}
			logrus.Error(err)

			return
		}

		err = c.rdbChat.PostToChannel(messagesEventNew, types.Message{
			Owner:   nickname,
			Content: string(bytes),
		})

		if err != nil {
			logrus.Error(err)

			return
		}

	}
}

func websocketWrite(conn *websocket.Conn, payload string) error {
	var msg types.Message

	err := json.NewDecoder(strings.NewReader(payload)).Decode(&msg)
	if err != nil {
		logrus.Info("Failed to decode message", err.Error())

		return err
	}

	wsResponse := fmt.Sprintf("> %s: %s", msg.Owner, msg.Content)
	err = conn.WriteMessage(websocket.TextMessage, []byte(wsResponse))
	if err != nil {
		logrus.Info("Failed to write message", err.Error())

		return err
	}

	return nil
}
