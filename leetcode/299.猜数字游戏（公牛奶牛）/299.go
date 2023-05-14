package main

import (
	"strconv"
)

func getHint(secret string, guess string) string {
	var A int
	var B int
	//var B = 0
	var cntS, cntG [10]int
	for i, _ := range secret {
		if secret[i] == guess[i] {
			A += 1
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		B += min(cntS[i], cntG[i])
	}

	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	println(getHint("1123157", "1157152"))
}
