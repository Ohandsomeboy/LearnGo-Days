package main

func equalPairs(grid [][]int) (ans int) {
	// 创建一个映射 cnt，将整数数组映射到它在 grid 中出现的次数。
	cnt := map[[200]int]int{}
	for _, row := range grid {
		// 创建一个包含行中的值的数组 a。
		a := [200]int{}
		for j, v := range row {
			a[j] = v
		}
		// 将 cnt[a] 加 1，以计算具有与行相同值集的行数。
		cnt[a]++
	}
	for j := range grid[0] {
		// 创建一个包含第 j 列中的值的数组 a。
		a := [200]int{}
		for i, row := range grid {
			a[i] = row[j]
		}
		// 将 cnt[a] 添加到 ans 中，以计算具有相同值集的行和列对的数量。
		ans += cnt[a]
	}
	return
}

func main() {

}
