package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websoket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool)
	}
}

func (s *server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conn[ws] = true

	s.readLoop(ws)
}

func (s *server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("Thank you for the msg"))
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(serverhandleWS))
	http.ListenAndServe(":3000", nil)

}