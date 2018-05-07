package chat

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func WsServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (s *Server) Listen() {
	for {
		select {
		case client := <-s.register:

			s.clients[client] = true

			fmt.Println("registration")
			fmt.Println(s.clients)
		case client := <-s.unregister:

			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
			}

			fmt.Println("unregistred: ")
			fmt.Println(s.clients)
		case message := <-s.broadcast:
			for client := range s.clients {

				ws, _ := client.ws.NextWriter(websocket.TextMessage)
				ws.Write(message)
			}
		}
	}
}

func (s *Server) BroadcastTest() {
	for {
		<-time.After(5 * time.Second)
		fmt.Println("start test broadcasting")
		s.broadcast <- getTestMessage()
	}
}

func getTestMessage() []byte  {
	return []byte("It's broadcast test")
}