/*
map: 无序键值对集合
2025.4.16 by dralee
*/
package main

import "fmt"

func main() {
	m1 := make(map[string]int)    // 空map
	m2 := make(map[string]int, 3) // 容量为3

	fmt.Printf("m1: %v, m2: %v\n", m1, m2)
}
