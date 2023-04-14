package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	//不对称
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

//时间复杂度：O(n)，其中n是字符串s的长度
//空间复杂度：O(n+∣Σ∣)，表示为栈的字符数量为O(n)，哈希表的空间为O(∣Σ∣)

func main() {
	var n = "({})[]"
	if isValid(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
