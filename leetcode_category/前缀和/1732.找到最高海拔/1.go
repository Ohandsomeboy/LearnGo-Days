package main

import "fmt"

func largestAltitude(gain []int) int {
	sum, max := 0, gain[0]
	for _, x := range gain {
		sum += x
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	m := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(m))
}
