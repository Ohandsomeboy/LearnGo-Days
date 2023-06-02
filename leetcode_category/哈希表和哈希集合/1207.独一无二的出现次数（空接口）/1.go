package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
	}
	t := map[int]struct{}{}
	for _, c := range m {
		t[c] = struct{}{}
	}
	return len(t) == len(m)
}

func main() {
	m := []int{1, 2, 2, 1, 1, 3, 2}
	fmt.Println(uniqueOccurrences(m))
}
