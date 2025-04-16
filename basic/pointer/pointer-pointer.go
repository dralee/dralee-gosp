/*
指向指针的指针变量
2025.4.16 by dralee
*/

package main
import "fmt"

func main(){
	var a int
	var ptr *int
	var pptr **int

	a = 2025
	ptr = &a  // 指向a的地址
	pptr = &ptr // 指向ptr的地址

	fmt.Printf("a = %d, *ptr = %d, **pptr = %d\n", a, *ptr, **pptr)

	a,b := 1,2
	swap(&a,&b)
	fmt.Printf("a = %d, b = %d\n", a, b)

}

func swap(x, y *int) {
	*x, *y = *y, *x
}