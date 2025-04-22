/*
主程
2025.4.22 by dralee
*/
package main

import (
	onlinenote "draleeonlinenote/basic"
	"draleeonlinenote/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := flag.String("host", "127.0.0.1", "http service host")
	port := flag.Int("port", 8012, "http service port")
	flag.Parse()

	s := onlinenote.NewServer()
	utils.NewLogger("draleeonlinenote")
	go s.Listen()
	http.HandleFunc("/", s.Home)
	http.HandleFunc("/ws", s.Accept)
	fmt.Println("start server")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
