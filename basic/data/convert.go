/*
类型转换
数值转换：strconv
类型转换：value.(type)或value.(T)
2025.4.17 by dralee
*/
package main

import (
	"fmt"
	"strconv"
)

func basic() {
	var str string = "123"
	var n int
	n, _ = strconv.Atoi(str)
	fmt.Printf("atoi: %d\n", n)

	str1 := strconv.Itoa(n)
	fmt.Printf("itoa: %s\n", str1)

	str2 := "3.1415926"
	f, _ := strconv.ParseFloat(str2, 64)
	fmt.Printf("parsefloat: %f\n", f)

	str3 := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("formatfloat: %s\n", str3)
}

func types() {
	var i interface{} = "Hello World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

// Writer接口
type Writer interface {
	Write([]byte) (int, error)
}

// 实现Writer接口
type StringWriter struct {
	str string
}

// 实现Write方法
func (w *StringWriter) Write(data []byte) (int, error) {
	w.str += string(data)
	return len(data), nil
}

func impTypes() {
	var w Writer = &StringWriter{} // 创建并赋给Writer接口变量

	sw := w.(*StringWriter) // 强制转换为StringWriter指针

	sw.str = "Hello World"
	fmt.Println(sw.str)

}

func print(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	default:
		fmt.Printf("unknown type\n")
	}
}

func printTest() {
	print(1)
	print("hello")
	print(123.5)
}

func main() {
	basic()

	types()

	impTypes()

	printTest()
}
