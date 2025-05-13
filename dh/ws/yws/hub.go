/*
管理房间
2025.5.13 by dralee
*/
package main

import (
	"sync"
)

type Hub struct {
	rooms map[string]*Room
	mu    sync.Mutex
}

var globalHub = &Hub{
	rooms: make(map[string]*Room),
}

func (h *Hub) GetRoom(id string) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()

	if room, ok := h.rooms[id]; ok {
		return room
	}
	room := NewRoom(id)
	h.rooms[id] = room
	return room
}
