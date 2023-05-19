package main

import (
	"fmt"
	"unicode"
)

func camelMatch(queries []string, pattern string) []bool {
	//n := len(queries)
	//res := make([]bool, n)
	//for i := 0; i < n; i++ {
	//	res[i] = true
	//	p := 0
	//	for _, c := range queries[i] {
	//		if p < len(pattern) && pattern[p] == byte(c) {
	//			p++
	//		} else if unicode.IsUpper(c) { // 判断这个字符是小写还是大写，大写返回true，小写返回fasle
	//			res[i] = false
	//			break
	//		}
	//	}
	//	if p < len(pattern) {
	//		res[i] = false
	//	}
	//}
	//return res

	n := len(queries)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = true
		p := 0
		for _, c := range queries[i] {
			if p < len(pattern) && pattern[p] == byte(c) {
				p++
			} else if unicode.IsUpper(c) {
				res[i] = false
				break
			}
		}
		if p < len(pattern) {
			res[i] = false
		}
	}
	return res
}

func main() {
	var queries = []string{"FooBar", "FooBarTest"}
	var s string = "FB"
	fmt.Println(camelMatch(queries, s))
}
