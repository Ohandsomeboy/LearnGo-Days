package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	t := []byte(s) // 将其转为byte类型，方便操作
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}

func main() {
	fmt.Println(reverseVowels("hello"))
}
