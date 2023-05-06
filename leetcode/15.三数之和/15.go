package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	if n < 3 {
		return ans
	}
	sort.Ints(nums) // 递增排序
	for i := 0; i < n; i++ {
		if nums[i] > 0 { // 起始数组大于0，后续也大于0.因为整个数组是有序的
			return ans
		}
		// 去重，当起始的值等于前一个元素，
		// 那么得到的结果将会和前一次相同.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := n - 1
		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				// 去重，因为 i 不变，
				// 当此时 l取的数的值与前一个数相同，所以不用在计算，直接跳
				for ; L < R && nums[L] == nums[L+1]; L = L + 1 {
				}
				//去重，因为 i不变，
				//当此时 r 取的数的值与前一个相同，所以不用在计算
				for ; L < R && nums[R] == nums[R-1]; R = R - 1 {
				}
				L = L + 1
				R = R - 1
				// 结果大于0，R指针左移
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R = R - 1
				// 结果小于0，L指针右移
			} else {
				L = L + 1
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
