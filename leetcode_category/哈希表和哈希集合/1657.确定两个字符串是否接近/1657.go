package main

func closeStrings(word1 string, word2 string) bool {
	m1 := map[rune]int{}
	m2 := map[rune]int{}
	if len(word1) != len(word2) {
		return false
	}
	for _, x := range word1 {
		m1[x]++
	}
	for _, x := range word2 {
		m2[x]++
	}
	// 数量判断
	if len(word1) != len(word2) {
		return false
	}
	// 种类判断
	for k, _ := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}
	// 对应字符串数量
	for key, _ := range m1 {
		for key2, _ := range m2 {
			if m1[key] == m2[key2] { // 有字符的数量是否相同
				m2[key2] = 0
				break
			}
		}
	}
	// 判断所有的字符是不是都有对应的数量的不同或相同的字符相等
	for _, v := range m2 {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	println(closeStrings("abc", "bca"))
}
