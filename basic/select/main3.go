/*
结束消息
*/
package main

import (
	"fmt"
	"time"
)

func Chann(ch chan int, chStop chan bool) {
	var i int
	i = 10
	for j := 0; j < i; j++ {
		ch <- j
		time.Sleep(time.Second * 1)
	}
	chStop <- true
}

func main() {
	ci := make(chan int)
	cb := make(chan bool)
	c :=0 

	go Chann(ci, cb)

	for {
		select {
		case c = <-ci:
			fmt.Println("received:", c)
			fmt.Println("channel")
		case s:= <-ci:
			fmt.Println("received:", s)
		case _ = <-cb:
			goto end
		}
	}
end:
	fmt.Println("end")
}