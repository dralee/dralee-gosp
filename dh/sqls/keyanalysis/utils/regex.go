/*
正则
2025.4.18 by dralee
*/
package utils

import (
	"regexp"
)

func Match(regex string, str string) bool {
	matched, _ := regexp.MatchString(regex, str)
	return matched
}

func MatchAll(regex string, str string) []string {
	return regexp.MustCompile(regex).FindAllString(str, -1)
}

func FindGroup(regex string, str string) []string {
	return regexp.MustCompile(regex).FindStringSubmatch(str)
}

func FindAllGroup(regex string, str string, index int) []string {
	groups := regexp.MustCompile(regex).FindAllStringSubmatch(str, -1)
	if groups == nil {
		return nil
	}

	if len(groups) == 0 {
		return nil
	}

	if index >= len(groups[0]) {
		panic("index out of range")
	}

	//fmt.Println("groups:", groups)
	result := make([]string, 0, len(groups))
	//fmt.Println("len:", len(groups))
	for _, g := range groups {
		// v := strings.TrimSpace(g[index])
		// if v == "" {
		// 	fmt.Println("v is space:", v)
		// }

		// if v == "" {
		// 	continue
		// }
		v := g[index]
		result = append(result, v)
	}
	return result
}

func FindAllGroupStr(regex string, str string) [][]string {
	return regexp.MustCompile(regex).FindAllStringSubmatch(str, -1)
}

func Replace(regex string, str string, replace string) string {
	return regexp.MustCompile(regex).ReplaceAllString(str, replace)
}
