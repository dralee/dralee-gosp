/*
指针
2025.4.16 by dralee
*/
package main

import "fmt"

func basic(){
	var a int = 10
	var b *int  // 声明指针变量

	fmt.Printf("empty ponter: %x, is nil: %d\n", b, b == nil)
	b = &a    // 指针变量指向a存储地址

	fmt.Printf("adress of a: %x\n", &a)
	fmt.Printf("address of b: %x\n", b)
	fmt.Printf("value of a: %d\n", a)
	fmt.Printf("value of b: %d\n", *b)
}

const ARRAY_SIZE = 3
func pointArray(){
	a := []int{1,2,3}
	var ptr [ARRAY_SIZE]*int // 声明一个指针数组

	for i := 0; i < ARRAY_SIZE; i++ {
		ptr[i] = &a[i]   // 指向数据元素中地址
	}

	for i := 0; i < ARRAY_SIZE; i++ {
		fmt.Printf("a[%d] = %d, *ptr[%d] = %d\n", i, a[i], i, *ptr[i])
	}

	*ptr[2] = 20  // 指向的是地址，因此改ptr[2]，同时会改变a[2]的值
	fmt.Printf("a[2] = %d, *ptr[2] = %d\n", a[2], *ptr[2])

}

func main(){
	basic()
	pointArray()


}