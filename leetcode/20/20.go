package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	// 检查括号序列长度是否为奇数
	if n%2 == 1 {
		return false
	}
	// 构建括号对映射关系
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 创建一个栈用于存放括号字符
	stack := []byte{}
	// 遍历括号序列中的每一个字符
	for i := 0; i < n; i++ {
		// 如果当前字符为右括号
		if pairs[s[i]] > 0 {
			// 栈为空或者栈顶元素与当前右括号所对应的左括号不匹配，返回false
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 当前字符为左括号，将其压入栈中
			stack = append(stack, s[i])
		}
	}
	// 栈为空则括号序列合法，否则不合法
	return len(stack) == 0
}

//时间复杂度：O(n)，其中n是字符串s的长度
//空间复杂度：O(n+∣Σ∣)，表示为栈的字符数量为O(n)，哈希表的空间为O(∣Σ∣)

func main() {
	var n = "({}[])"
	if isValid(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
