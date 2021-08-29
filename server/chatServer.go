package server

import (
	"github.com/PhilipHunkevych/go-messaging-app/pkg/types"
	"log"
)

type ChatServer struct {
	channels *types.ChatChannels
}

func NewChatServer(c *types.ChatChannels) *ChatServer {
	return &ChatServer{
		channels: c,
	}
}

func (cs *ChatServer) Start() {
	clients := make(map[string]chan *types.Msg)

	for {
		select {
		case req := <-cs.channels.ClientRequests:
			clients[req.Nickname] = req.MsgChan
			log.Println("Websocket connected: " + req.Nickname)
		case ClientNick := <-cs.channels.ClientDisconnects:
			close(clients[ClientNick])
			delete(clients, ClientNick)
			log.Println("Websocket disconnected: " + ClientNick)
		case msg := <-cs.channels.Messages:
			for _, msgChan := range clients {
				if len(msgChan) < cap(msgChan) {
					msgChan <- msg
				}
			}
		}
	}
}
