/*
集合工具
2025.4.18 by dralee
*/
package utils

func Contains[T comparable](arr []T, e T) bool {
	for _, a := range arr {
		if a == e {
			return true
		}
	}
	return false
}
