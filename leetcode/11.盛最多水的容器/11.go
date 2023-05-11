package _1_盛最多水的容器

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
