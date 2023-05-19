package main

import (
	"fmt"
	"sort"
)

func meticSubarrays(nums []int) bool {
	a := make([]int, len(nums))
	copy(a, nums)
	sort.Ints(a)
	ch := a[1] - a[0]
	for i := 0; i < len(a)-1; i++ {
		if a[i+1]-a[i] != ch {
			return false
		}
	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) (c []bool) {
	diff := []int{}

	for j := 0; j < len(l); j++ {
		diff = nums[l[j] : r[j]+1]
		//fmt.Println(l[j], r[j], diff)

		if meticSubarrays(diff) {
			c = append(c, true)
		} else {
			c = append(c, false)
		}
		diff = []int{}
	}
	return c
}

func main() {
	var nums = []int{4, 6, 5, 9, 3, 7}
	var l = []int{0, 0, 2}
	var r = []int{2, 3, 5}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}
