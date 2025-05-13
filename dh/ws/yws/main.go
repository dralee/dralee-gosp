/*
主程
2025.5.13 by dralee
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws/", handleWebSocket)
	http.HandleFunc("/", handleHome)
	log.Println("Yjs WebSocket gateway started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
