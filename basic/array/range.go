/*
range: 可对array,slice,map进行迭代,key,value
2025.4.16 by dralee
*/
package main

import (
	"fmt"
)

func array() {
	num := [3]int{1, 2, 3}
	for i, v := range num {
		println(i, v)
	}
	for key := range num {
		println(key) // 索引
	}
	for _, v := range num {
		println(v) // 值
	}

	for i, c := range "golang" {
		fmt.Printf("index %d is %c\n", i, c)
	}
}

func mapRange() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m["d"] = 40
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func channelRange() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	array()
	mapRange()
	channelRange()
}
