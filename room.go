package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan []byte
	clients map[*client]bool
	join    chan *client
	leave   chan *client
}

func (r *room) run() {
	select {
	case msg := <-r.forward:
		for client := range r.clients {
			client.to <- msg
		}
	case client := <-r.join:
		r.clients[client] = true
	case client := <-r.leave:
		delete(r.clients, client)
		close(client.to)
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err.Error())
	}

	client := &client{
		room:   r,
		to:     make(chan []byte, messageBufferSize),
		socket: socket,
	}

	r.join <- client
	defer func() {
		r.leave <- client
	}()

	go client.SendMessage()
	client.ReadMessage()
}

func NewRoom() *room {
	return &room{
		forward: make(chan []byte),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}