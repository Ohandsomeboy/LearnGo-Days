package main

func minOperations(nums []int) (count int) {
	if len(nums) == 1 {
		return 0
	}
	pre := nums[0] - 1
	for _, x := range nums {
		if pre+1 > x {
			pre = pre + 1
		} else {
			pre = x
		}
		count += pre - x
	}
	return
}

// func minOperations(nums []int) (ans int) {
//     pre := nums[0] - 1
//     for _, x := range nums {
//         pre = max(pre+1, x)
//         ans += pre - x
//     }
//     return
// }

// func max(a, b int) int {
//     if b > a {
//         return b
//     }
//     return a
// }

func main() {
	var num_array = []int{1, 5, 2, 4, 1}
	println(minOperations(num_array))
}
