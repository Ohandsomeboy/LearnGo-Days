package main

// https://leetcode.cn/problems/majority-element/solution/duo-shu-yuan-su-by-leetcode-solution/
// 排序，随机化，分治，Boyer-Moore投票，
// 哈希表
func majorityElement(nums []int) int {
	n := len(nums)
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for i, v := range m {
		if v > n/2 {
			return i
		}
	}
	return n
}
