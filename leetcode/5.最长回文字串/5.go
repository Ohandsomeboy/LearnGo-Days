package main

import "fmt"

// 暴力解法
//func reverseString(s string) bool {
//	runes := []rune(s)
//
//	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
//		runes[from], runes[to] = runes[to], runes[from]
//	}
//	if s == string(runes) {
//		return true
//	}
//
//	return false
//}
//
//func longestPalindrome(s string) string {
//	var maxR string
//	var max = 0
//	for i := 0; i <= len(s)-1; i++ {
//		for j := i + 1; j <= len(s); j++ {
//			if reverseString(s[i:j]) && len(s[i:j]) > max {
//				maxR = s[i:j]
//				max = len(s[i:j])
//			}
//		}
//	}
//	return maxR
//}

// 动态规划
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1]
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for length := 2; length <= len(s); length++ {
		for start := 0; start < len(s)-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] {
				continue
			} else if length < 3 {
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1]
			}
			if dp[start][end] && (end-start+1) > len(result) {
				result = s[start : end+1]
			}
		}
	}
	return result
}

func main() {
	var s = "ibvjkmpyzsifuxcabqqpahjdeuzaybqsrsmbfplxycsafogotliyvhxjtkrbzqxlyfwujzhkdafhebvsdhkkdbhlhmaoxmbkqiwiusngkbdhlvxdyvnjrzvxmukvdfobzlmvnbnilnsyrgoygfdzjlymhprcpxsnxpcafctikxxybcusgjwmfklkffehbvlhvxfiddznwumxosomfbgxoruoqrhezgsgidgcfzbtdftjxeahriirqgxbhicoxavquhbkaomrroghdnfkknyigsluqebaqrtcwgmlnvmxoagisdmsokeznjsnwpxygjjptvyjjkbmkxvlivinmpnpxgmmorkasebngirckqcawgevljplkkgextudqaodwqmfljljhrujoerycoojwwgtklypicgkyaboqjfivbeqdlonxeidgxsyzugkntoevwfuxovazcyayvwbcqswzhytlmtmrtwpikgacnpkbwgfmpavzyjoxughwhvlsxsgttbcyrlkaarngeoaldsdtjncivhcfsaohmdhgbwkuemcembmlwbwquxfaiukoqvzmgoeppieztdacvwngbkcxknbytvztodbfnjhbtwpjlzuajnlzfmmujhcggpdcwdquutdiubgcvnxvgspmfumeqrofewynizvynavjzkbpkuxxvkjujectdyfwygnfsukvzflcuxxzvxzravzznpxttduajhbsyiywpqunnarabcroljwcbdydagachbobkcvudkoddldaucwruobfylfhyvjuynjrosxczgjwudpxaqwnboxgxybnngxxhibesiaxkicinikzzmonftqkcudlzfzutplbycejmkpxcygsafzkgudy"
	//println(s[0:5])
	//println(len(s))
	fmt.Println(longestPalindrome(s))
}
