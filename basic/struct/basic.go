/*
结构体
定义结构体需要使用type和struct，struct定义新数据类型，type定义新类型
2025.4.16 by dralee
*/
package main

import "fmt"

type Person struct {
	name string
	age int
	address string
}

func main(){
	p := Person{"Bob", 20, "New York"}
	fmt.Println(p)
	fmt.Println("info: ", p.info())

	fmt.Printf("adress of p: %p\n", &p) // adress of p: 0xc0000b20f0
	hello(p)
	hello2(&p)
}

func (p Person) info() string {
	return fmt.Sprintf("Name: %s, Age: %d, Address: %s", p.name, p.age, p.address)
}

func hello(p Person){ // 传递是拷贝副本
	fmt.Printf("adress of p: %p\n", &p) // adress of p: 0xc0000b2150
	fmt.Println("hello", p.name)
}

func hello2(p *Person){
	fmt.Printf("adress of p: %p\n", p) // adress of p: 0xc0000b20f0
	fmt.Println("hello for", p.name)
}