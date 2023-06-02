package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}

func main() {
	m := []int{7, 7, 3, 2, 6, 5, 6}
	fmt.Println(pivotIndex(m))
}
