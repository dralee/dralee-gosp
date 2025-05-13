/*
客户端连接处理
2025.5.13 by dralee
*/

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 🔒 JWT 校验
	log.Println("the ws handler....")
	userID, err := ValidateJWT(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		conn.Close()
		return
	}

	room := globalHub.GetRoom(roomID)
	client := &Client{
		conn:   conn,
		send:   make(chan []byte, 256),
		room:   room,
		userID: userID,
	}

	room.register <- client

	go client.writePump()
	client.readPump()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}
