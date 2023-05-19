package main

func carPooling(trips [][]int, capacity int) bool {
	n := 1001
	diff := []int{}
	for i := 0; i < n; i++ {
		diff = append(diff, 0)
	}
	// 建立查分数组
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}

	// 计算前缀和
	var preSum int = 0
	for i := 0; i < n; i++ {
		preSum += diff[i]
		if preSum > capacity {
			return false
		}
	}
	return true
}

func main() {
	var num_array = [][]int{{2, 1, 5}, {3, 3, 7}}
	println(carPooling(num_array, 4))
}
