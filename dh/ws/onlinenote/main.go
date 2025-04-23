/*
主程
2025.4.22 by dralee
*/
package main

import (
	"draleeonlinenote/basic"
	"draleeonlinenote/session"
	"draleeonlinenote/utils"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	host := flag.String("host", "127.0.0.1", "http service host")
	port := flag.Int("port", 8012, "http service port")
	flag.Parse()

	dsn := "root:1234@tcp(127.0.0.1:3306)/notedb?charset=utf8mb4&parseTime=True&loc=Local"
	s := session.NewServer(dsn)
	utils.NewLogger("draleeonlinenote")
	go s.Listen()
	http.HandleFunc("/", s.Home)
	http.HandleFunc("/ws", s.Accept)
	basic.Info("start server")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
	if err != nil {
		basic.Errorf("ListenAndServe: %v", err)
	}
}
