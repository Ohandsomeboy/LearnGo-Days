package main

import "fmt"

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, k := 0, 0
	for i < n && k < m {
		if s[i] == t[k] {
			i++
		}
		k++
	}
	return i == n
}

func main() {
	var s, t string
	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
