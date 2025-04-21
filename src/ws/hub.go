package ws

import "github.com/gorilla/websocket"

type Hub struct {
	// all the registered clients
	clients map[*Client]bool

	// the inbound messages from the clients
	broadcast chan []byte

	// register requests from the clients
	register chan *Client

	// unregister requests from the clients.
	unregister chan *Client
}

type Client struct {
	conn *websocket.Conn
	hub  *Hub
	send chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {

		case client := <-hub.register:
			hub.clients[client] = true

		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}

		case message := <-hub.broadcast:
			for peer := range hub.clients {
				select {
				case peer.send <- message:

				default:
					close(peer.send)
					delete(hub.clients, peer)
				}
			}
		}
	}
}
