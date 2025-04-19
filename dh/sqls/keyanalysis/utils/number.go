/*
数值转换
2025.4.19 by dralee
*/
package utils

import "strconv"

func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func ToUInt64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
