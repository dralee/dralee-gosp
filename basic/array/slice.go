/*
由于数组是不可变长度的
切片: 动态数组,可以追加元素,可使切片容量增大
定义: make([]type, len)
var slice1 []type = make([]type, len)
slice1 := make([]type, len)
直接初始: s := [] int{1,2,3}, s := arr[:] 从下标到结束
2025.4.16 by dralee
*/
package main

import "fmt"

func truncateSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	showSlice(s)

	s1 := s[1:4] // 从下标1到下标4（不包含）
	showSlice(s1)

	s2 := s[:3] // 从下标0到下标3（不包含）
	showSlice(s2)

	s3 := s[4:] // 从下标4到结束
	showSlice(s3)

	s3 = append(s3, 10, 20, 30) // 将10,20,30追加到s3
	showSlice(s3)

	s4 := make([]int, len(s3), cap(s3))
	copy(s4, s3) // 将s3拷贝到s4
	showSlice(s4)

}

func main() {
	var arr = make([]int, 3, 5) // len:3, capacity: 5
	var empty []int             // 空切片

	fmt.Println("slice is empty:", empty == nil)
	fmt.Println(arr)

	appendSlice(arr)

	truncateSlice()
}

func showSlice(s []int) {
	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s), cap(s), s)
}

func appendSlice(s []int) []int {
	s = append(s, 4, 5) // 新增2个元素
	fmt.Println(s)
	return s
}
