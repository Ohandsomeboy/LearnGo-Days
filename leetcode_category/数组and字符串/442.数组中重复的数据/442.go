package main

import "fmt"

func findDuplicates(nums []int) (ans []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v >= 2*n+1 {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	var num_array = []int{2, 2}
	fmt.Println(findDuplicates(num_array))
}
