/*
没消息则使用default
*/
package main

import (
	"fmt"
)

func main() {
	// 定义两个通道
	c1 := make(chan string)
	c2 := make(chan string)

	// 两个goroutine
	go func() {
		for {
			c1 <- "from 1"
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
		}
	}()

	// 使用select语句非阻塞地从两个通道接收数据
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		default:
			fmt.Println("no message received")
		}
	}
}