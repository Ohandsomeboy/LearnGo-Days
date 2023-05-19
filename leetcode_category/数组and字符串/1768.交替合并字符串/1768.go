package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var s strings.Builder
	n1 := len(word1)
	n2 := len(word2)
	var min int
	if n1 > n2 {
		min = n2

	} else if n2 > n1 {
		min = n1
	} else if n1 == n2 {
		min = n1
	}
	i := 0
	for ; i < min; i++ {
		s.WriteRune(rune(word1[i]))
		s.WriteRune(rune(word2[i]))
	}
	if n1 > n2 {
		s.WriteString(word1[i:])
	} else {
		s.WriteString(word2[i:])
	}
	str := s.String()
	return str
}

func main() {
	var s1 string = "abc"
	var s2 string = "pqr"
	fmt.Println(mergeAlternately(s1, s2))
}
