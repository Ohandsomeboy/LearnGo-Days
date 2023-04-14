package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums) // 排序，从小到大
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + nums[i]
	}
	ans := []int{}
	for _, q := range queries {
		// x找不到时就会返回比x大的数的index，
		// 如果 x 不存在，则可能返回 ，切片长度len(a)
		println(f[0], f[1], f[2], f[3])
		idx := sort.SearchInts(f, q+1)
		ans = append(ans, idx)
	}
	return ans
}

func main() {
	var num_array = []int{1, 4, 5, 2}
	var q = []int{3, 10, 6}
	fmt.Println(answerQueries(num_array, q))
}
