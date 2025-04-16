/*
数组
数据的大小是类型的一部分，因此不同大小的数组是不兼容的，即[5]int与[10]int是不同类型的
不确定长度的，可使用...代替长度，[...]int代表任何长度的int切片
2025.4.16 by dralee
*/
package main
import (
	"fmt"
)

func intArray(){
	fmt.Println("intArray:")

	var n [10]int // 10个int
	fmt.Println(n)
	fmt.Println(n[2])
	n[2] = 12
	fmt.Println(n)

	var i int
	for i = 0; i < 10; i++ {
		n[i] = i + 10
	}

	for i = 0; i < 10; i++ {
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}
}

func floatArray(){
	fmt.Println("floatArray:")
	var i int

	balance := [5]float32{100.0, 200.0, 300.0, 400.0, 500.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("Balance[%d] = %f, ", i, balance[i])
	}
	fmt.Println()

	balance2 := [...]float32{100.0, 200.0, 300.0, 400.0, 500.0}
	for i = 0; i < 5; i++ {
		fmt.Printf("Balance[%d] = %f, ", i, balance2[i])
	}
	fmt.Println()

	balance3 := [5]float32{1:100.0, 3:200.0} // 只初始化索引为1和3的元素
	for i = 0; i < 5; i++ {
		fmt.Printf("Balance[%d] = %f, ", i, balance3[i])
	}
	fmt.Println()
}

func multidimensionalArray(){
	fmt.Println("multidimensionalArray:")
	// 多维数组
	values := [][]int{}

	row1 := []int{1,2,3}
	row2 := []int{4,5,6}
	values = append(values, row1) // 通过append将row1添加到values
	values = append(values, row2)
	fmt.Println(values)

	fmt.Printf("first element: %d", values[0][0])

	a := [3][4]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12}, // 必须添加逗号，不然就需要将}顶上来
	}

	fmt.Println(a)
}

func average(values []float32) float32 {
	var sum float32 = 0
	for i:=0; i < len(values); i++ {
		sum += values[i]
	}

	return sum / float32(len(values))
}

func testAverage(){
	fmt.Println("testAverage:")
	values := []float32{1.5,2.3,3.2,4.1,5.0}
	fmt.Println(average(values))

}

func main(){
	intArray()

	floatArray()

	multidimensionalArray()

	testAverage()
}