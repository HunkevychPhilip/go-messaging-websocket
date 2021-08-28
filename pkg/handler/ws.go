package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func home(rw http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(rw, "Home Page")
}

func newWS(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err.Error())

		return
	}
	log.Println("Client successfully connected.")

	go writer(conn)
	//go reader(conn)
}

func writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(time.Second * 1)

		i := 0
		for c := range ticker.C {
			fmt.Printf("Ticker:  %+v\n", c)

			msg := fmt.Sprintf("Writing message number %d.", i)
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("Failed write message" + err.Error())

				return
			}
			i++
		}
	}
}

func newReader(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err.Error())

		return
	}
	log.Println("Client successfully connected.")

	go reader(conn)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())

			return
		}

		log.Println(string(p))
	}
}
