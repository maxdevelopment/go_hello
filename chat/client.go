package chat

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"fmt"
)

type Client struct {
	ws             *websocket.Conn
	server         *Server
	messageChannel chan []byte
}

func StartWs(s *Server, writer http.ResponseWriter, request *http.Request) {
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal(err)
	}
	messageChannel := make(chan []byte)
	client := &Client{ws: ws, server: s, messageChannel: messageChannel}
	s.register <- client

	go client.readMsg()
}

func (c *Client) readMsg() {
	defer func() {
		fmt.Println("deffer")
		c.server.unregister <- c
		c.ws.Close()
	}()

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("MSG: ", message)
		//select {
		//case message:= <-c.messageChannel:
		//	fmt.Println("MSG: ", message)
		//
		//}
	}
}
