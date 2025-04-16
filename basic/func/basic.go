/*
基本
2025.4.16 by dralee
*/

package main
import (
	"fmt"
	"math"
)

func swap_str(x, y string) (string, string) {
	return y, x
}

func swap(x , y *int) {
	*x, *y = *y, *x
}

// 闭包
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	s1, s2 := "world", "hello"
	s3,s4 := swap_str(s1, s2)

	fmt.Println(s3, s4)

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)


	fmt.Println("闭包")
	nextNumber := getSequence()
	// 调用nextNumber() i 变量自增1，并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	// 调用nextNumber1() i 变量自增1，并返回，新变量，则新从0开始
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}