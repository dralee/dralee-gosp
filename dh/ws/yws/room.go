/*
房间定义
2025.5.13 by dralee
*/
package main

type Room struct {
	id         string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewRoom(id string) *Room {
	r := &Room{
		id:         id,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
	go r.run()
	return r
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
		case client := <-r.unregister:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.broadcast:
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}
