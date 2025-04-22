/*
客户端
2025.4.22 by dralee
*/
package onlinenote

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 消息类型
	MsgTypeText    = "text"
	MsgTypeImage   = "image"
	writeWait      = 10 * time.Second    // time to write a message to the peer
	pongWait       = 60 * time.Second    // time to read the next pong message from the peer
	pingPeriod     = (pongWait * 9) / 10 // send pings to peer with this period
	maxMessageSize = 512                 // max message size
)

type Client struct {
	server OnlineServer
	conn   *websocket.Conn
	send   chan []byte
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(Newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) read() {
	defer func() {
		c.server.offline(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				Errorf("error: %v", err)
			}
			break
		}
		message = append(message, Newline...)
		c.server.Broadcast(message)
	}
}
