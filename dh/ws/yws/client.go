/*
客户端定义
2025.5.13 by dralee
*/
package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	room   *Room
	userID string
}

func (c *Client) readPump() {
	defer func() {
		c.room.unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		// 将数据广播给房间内所有人
		c.room.broadcast <- msg
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		c.conn.WriteMessage(websocket.BinaryMessage, msg)
	}
}
