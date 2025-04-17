/*
map: 无序键值对集合
2025.4.16 by dralee
*/
package main

import "fmt"

func demoMap() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	a := m["a"]
	b, ok := m["b"]
	d, ok := m["d"]
	fmt.Printf("len: %d\n", len(m))
	fmt.Printf("a:%d\nb:%d, exists:%v\n", a, b, ok)
	fmt.Printf("d:%d, exists:%v\n", d, ok)

	for k, v := range m {
		fmt.Printf("%s=%d ", k, v)
	}
	fmt.Println()

	delete(m, "b")
	fmt.Println("after delete b:", m)

}

func main() {
	m1 := make(map[string]int)    // 空map
	m2 := make(map[string]int, 3) // 容量为3

	fmt.Printf("m1: %v, m2: %v\n", m1, m2)

	demoMap()
}
