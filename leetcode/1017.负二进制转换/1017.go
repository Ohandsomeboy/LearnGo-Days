package main

import "fmt"

func baseNeg2(n int) (s string) {
	if n == 0 {
		return "0"
	}
	ans := []byte{}
	k := 1
	for n != 0 {
		if n%2 != 0 {
			ans = append(ans, '1')
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}

	// 反转字符串
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

func main() {
	fmt.Printf(baseNeg2(4))
}
