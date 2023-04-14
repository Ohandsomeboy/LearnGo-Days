package main

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	count := 0
	num := arr[0]
	for i := 1; i < n; i++ {
		if i < n-1 {
			if num <= arr[i] && arr[i] <= arr[i+1] {
				num = arr[i]
			} else {
				count += 1
			}
		} else if i == n-1 {
			if num < arr[i] {
				num = arr[i]
			} else {
				count += 1
			}
		}
		if i+1 == n {
			break
		}
	}
	return count
}

func main() {
	var array = []int{5, 4, 3, 2, 1}
	println(findLengthOfShortestSubarray(array))
}
