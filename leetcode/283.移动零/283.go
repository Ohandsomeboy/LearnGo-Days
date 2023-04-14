package main

import "fmt"

func moveZeroes(nums []int) []int {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return nums
}

func main() {
	var a = []int{0, 0, 1, 57, 8, 0}
	fmt.Println(moveZeroes(a))
}
