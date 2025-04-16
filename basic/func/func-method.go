/*
函数方法
2025.4.16 by dralee
*/

package main
import (
	"fmt"
	"math"
)

// 定义结构体
type Circle struct {
	radius float64
}

func main(){
	var c1 Circle
	c1.radius = 12.5
	fmt.Println("The area of the circle is", c1.area())
	fmt.Println("The circumference of the circle is", c1.circumference())
}

// Circle结构体类型的方法(method)
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Circle结构体类型的方法(method)
func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}