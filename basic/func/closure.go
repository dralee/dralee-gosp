/*
闭包场景
2025.4.16 by dralee
*/
package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}

	// 调用匿名函数
	fmt.Println(add(1, 2))

	multiply := func(x, y int) int {
		return x * y
	}

	// 调用匿名函数
	fmt.Println("3 * 4 =",multiply(3, 4))

	// 将匿名函数作为参数传递给其他函数
	calc := func(operator func(int, int) int, x, y int) int {
		return operator(x, y)
	}

	sum := calc(add, 1, 2)
	fmt.Println("1 + 2 =", sum)
	product := calc(multiply, 3, 4)
	fmt.Println("3 * 4 =", product)

	diff := calc(func(a,b int) int {
		return a - b
	}, 8, 3)
	fmt.Println("8 - 3 =", diff)
}