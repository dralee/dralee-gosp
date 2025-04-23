/*
服务端

2025.4.22 by dralee
*/
package session

import (
	"draleeonlinenote/basic"
	"draleeonlinenote/note"
	"draleeonlinenote/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// 服务接口
type OnlineServer interface {
	Listen()                                            // 启动监听
	Home(w http.ResponseWriter, r *http.Request)        // 主页
	Accept(w http.ResponseWriter, r *http.Request)      // 接受连接
	offline(c *Client)                                  // 下线
	Broadcast(message []byte)                           // 广播
	Login(w http.ResponseWriter, r *http.Request) error // 登录
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	clients     map[*Client]bool // registered clients
	broadcast   chan []byte      // broadcast channel
	register    chan *Client     // register requests
	unregister  chan *Client     // unregister requests
	userService user.UserService // 用户服务
	noteService note.NoteService // 笔记服务
}

func NewServer(dsn string) *Server {
	repo, err := user.NewDefaultUserRepository(dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	repoNote, _ := note.NewDefaultNoteRepository(dsn)
	userService := user.NewDefaultUserService(repo)
	noteService := note.NewDefaultNoteService(repoNote)
	return &Server{
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		userService: userService,
		noteService: noteService,
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
	id, name, err := s.userService.TokenInfo(w, r)
	if err != nil {
		basic.Error(err)
		return
	}
	client := &Client{server: s, conn: ws, send: make(chan []byte, 256), userId: id, userName: name}
	s.register <- client
	go client.write()
	go client.read()
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) error {
	_, err := s.userService.Login(w, r)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	isStatic := strings.HasPrefix(r.URL.Path, "/lib/")
	basic.Info("method: %s, path: %s, %v", r.Method, r.URL.Path, isStatic)

	if r.URL.Path != "/" && r.URL.Path != "/login" && r.URL.Path != "/logout" && !isStatic {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" && r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "GET" {
		if !isStatic {
			s.HomePage(w, r)
			return
		}

		// if r.URL.Path == "/login" {
		// 	http.ServeFile(w, r, "./web/login.html")
		// 	return
		// }
		if strings.HasPrefix(r.URL.Path, "/lib/") {
			file := fmt.Sprintf("./web%s", r.URL.Path)
			basic.Info("file: %s", file)
			http.ServeFile(w, r, file)
			return
		}
	}

	if r.Method == "POST" {
		if r.URL.Path == "/login" {
			err := s.Login(w, r)
			if err != nil {
				w.Write(basic.ToJson(basic.Fail(401, err.Error())))
				return
			}
			//http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if r.URL.Path == "/logout" {
			s.userService.Logout(w, r)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

}

func (s *Server) HomePage(w http.ResponseWriter, r *http.Request) bool {
	token, err := r.Cookie(basic.TokenKey)
	basic.Info("token: %v", token)

	if err != nil {
		http.ServeFile(w, r, "./web/login.html")
		return false
	}
	err = s.userService.Validate(token.Value)
	if err != nil {
		http.ServeFile(w, r, "./web/login.html")
		return false
	}

	http.ServeFile(w, r, "./web/index.html")
	return true
}

func (s *Server) Broadcast(message []byte) {
	s.broadcast <- message
}

func (s *Server) offline(c *Client) {
	s.unregister <- c
}
