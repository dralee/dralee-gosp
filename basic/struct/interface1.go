/*
接口
2025.4.17 by dralee
*/
package main

import "fmt"

type Mobile interface {
	Call()
}

type Android struct{}
type Iphone struct{}

func (a Android) Call() {
	fmt.Println("Android call")
}
func (i Iphone) Call() {
	fmt.Println("Iphone call")
}

func main() {
	var m Mobile

	fmt.Println("method 1:")
	m = Android{}
	m.Call()
	m = Iphone{}
	m.Call()

	fmt.Println("method 2:")
	var m1 Mobile
	m1 = new(Android)
	m1.Call()
	m1 = new(Iphone)
	m1.Call()
}
