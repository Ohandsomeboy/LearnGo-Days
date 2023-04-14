package main

import "fmt"

func isPalindrome(x int) bool {
	var v int
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	num := 0
	for {
		v = x % 10
		num = num*10 + v
		x = x / 10
		if x <= num {
			break // 跳出for循环
		}
	}
	return num/10 == x || num == x
}

func main() {
	var c int = 100
	fmt.Println(isPalindrome(c))
}
