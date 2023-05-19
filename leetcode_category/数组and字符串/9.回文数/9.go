package main

import "fmt"

// 反转了一半的数字，如果是回文数则应该相等。
// 注意点在于回文数 奇数 或 偶数 的情况会有所不同。
// 奇数时候需要去掉一位（除以10并取整）
func isPalindrome(x int) bool {
	var v int
	if x < 0 || (x%10 == 0 && x != 0) { // 1.排除小于0的数 2.回文数是和%10不等于0的 3.排除x==0
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
	var c int = 1001
	fmt.Println(isPalindrome(c))
}
