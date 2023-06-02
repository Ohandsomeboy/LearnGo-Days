package main

import "fmt"

func removeStars(s string) string {
	m := make([]byte, 0)
	c := 0
	for i, _ := range s {
		if s[i] != '*' {
			m = append(m, s[i])
			c++
		} else if s[i] == '*' {
			m = append(m[:c-1], m[c:]...)
			c--
		}
	}
	s2 := string(m)
	return s2
}

func main() {
	v := "leet**cod*e"
	fmt.Println(removeStars(v))
}
