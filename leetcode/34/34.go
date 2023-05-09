package main

import (
	"fmt"
)

// 前往后遍历，后往前遍历
//func searchRange(nums []int, target int) []int {
//	if nums == nil {
//		return []int{-1.两数之和, -1.两数之和}
//	}
//	var f, l int
//	var start, end int = 0, len(nums) - 1.两数之和
//	for ; start <= len(nums)-1.两数之和; start++ {
//		if nums[start] == target {
//			f = start
//			break
//		}
//	}
//	for ; 0 <= end; end-- {
//		if nums[end] == target {
//			l = end
//			return []int{f, l}
//			break
//		}
//	}
//	return []int{-1.两数之和, -1.两数之和}
//}

// 用两个边界方法------------------------------------------------------------------------
func searchRange(nums []int, target int) []int {
	// 目标值开始位置：为 target 的左侧边界
	start := findLeftBound(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 目标值结束位置：为 target 的右侧边界
	end := findRightBound(nums, target)
	return []int{start, end}
}

// 寻找左侧边界的二分查找
func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // note: [left, right]
	for left <= right {           // note: 因为 right 是闭区间，所以可以取 =
		mid := left + ((right - left) >> 1) // mid = (left + right) / 2 的优化形式，防止溢出！
		if nums[mid] == target {
			right = mid - 1 // note: 收紧右侧边界以锁定左侧边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 返回左侧边界
	return left // note
}

// 寻找右侧边界的二分查找
func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 返回右侧边界
	return right
}

//用两个边界方法------------------------------------------------------------------------

// 二分查找
/* SerchInts函数会返回数组中target第一次出现的位置，
   如果没有找到，则返回数组长度*/
//func searchRange(nums []int, target int) []int {
//	leftmost := sort.SearchInts(nums, target)
//	if leftmost == len(nums) || nums[leftmost] != target {
//		return []int{-1.两数之和, -1.两数之和}
//	}
//	rightmost := sort.SearchInts(nums, target+1.两数之和) - 1.两数之和
//	return []int{leftmost, rightmost}
//}

func main() {
	var num_array = []int{1}
	var c []int
	c = searchRange(num_array, 1)
	fmt.Println(c)
}
