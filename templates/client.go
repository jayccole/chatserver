package main

import (
	"github.com/gorilla/websocket"
)

// client for single user

type client struct {

	//web socket for the client
	socket *websocket.Conn

	//recieve from other clients
	receive chan []byte

	//room for chatting
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
