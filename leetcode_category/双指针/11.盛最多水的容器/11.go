package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	n := len(height)
	i, j, ans := 0, n-1, 0
	for i < j {
		if height[i] < height[j] {
			ans = max(height[i]*(j-i), ans)
			i++
		} else {
			ans = max(height[j]*(j-i), ans)
			j--
		}
	}
	return ans
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
