/*
服务端

2025.4.22 by dralee
*/
package onlinenote

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 服务接口
type OnlineServer interface {
	Listen()                                       // 启动监听
	Home(w http.ResponseWriter, r *http.Request)   // 主页
	Accept(w http.ResponseWriter, r *http.Request) // 接受连接
	offline(c *Client)                             // 下线
	Broadcast(message []byte)                      // 广播
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	clients    map[*Client]bool // registered clients
	broadcast  chan []byte      // broadcast channel
	register   chan *Client     // register requests
	unregister chan *Client     // unregister requests
}

func NewServer() *Server {
	return &Server{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (s *Server) Listen() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
		case message := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}
}

func (s *Server) Accept(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{server: s, conn: ws, send: make(chan []byte, 256)}
	s.register <- client
	go client.write()
	go client.read()
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./web/index.html")
}

func (s *Server) Broadcast(message []byte) {
	s.broadcast <- message
}

func (s *Server) offline(c *Client) {
	s.unregister <- c
}
