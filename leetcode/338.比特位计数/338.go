package main

import "fmt"

func countBits(n int) []int {
	// make（切片类型，长度，容量）
	a := make([]int, 0, 0)
	//a := []int{}
	//var a []int
	count := 0
	for i := 0; i < n+1; i++ {
		j := i
		for j != 0 {
			j = j & (j - 1)
			count++
		}
		a = append(a, count)
		count = 0
	}
	return a
}

func main() {
	fmt.Println(countBits(5))
}
