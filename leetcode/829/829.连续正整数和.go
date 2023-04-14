package main

func consecutiveNumbersSum(n int) int {
	seem := make(map[int]int)
	sum := 0
	for i := 1; i < n+1; i++ {
		//println(i)
		for j := i; j < n+1; j++ {
			sum += j
			seem[sum] += 1
		}
		sum = 0
	}
	return seem[n]
}

func main() {
	println(consecutiveNumbersSum(88888888))
}
