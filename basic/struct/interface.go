/*
接口
没有关键字显式声明某个类型实现了某个接口；只要一个类型实现了接口要求的所有方法，那么这个类型就实现了这个接口。
接口变量可存储任何实现了该接口的值。
接口变量两部分：动态类型(dynamic type)和动态值(dynamic value)
零值接口：接口变量的零值是nil，一个未初始化的接口变量的零值是nil，且不包含任何值（包括动态类型和动态值）
空接口：空接口可以存储任何类型的值, interface{}
接口常见方法：多态（不同类型实现同一接口，实现多态行为），解耦（通过接口定义依赖关系，降低模块之间的耦合），泛化（使用空接口interface{}表示任意类型）
2025.4.17 by dralee
*/
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义圆
type Circle struct {
	radius float64
}

// 实现Shape接口
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Reader interface {
	Read() string
}
type Writer interface {
	Write(data string)
}

// 接口组合
type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

// 实现Read,Write
func (f File) Read() string {
	return "read data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}

func dynamicValue() {
	var i interface{} = 1
	fmt.Println("i:", i)
	i = "hello"
	fmt.Println("i:", i)

	var i1 interface{}
	fmt.Println("i1:", i1, "is nil: ", i1 == nil)
}

func main() {
	c := Circle{radius: 12.5}
	fmt.Println("The area of the circle is", c.Area())
	fmt.Println("The circumference of the circle is", c.Perimeter())

	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("hello world")

	dynamicValue()
}
