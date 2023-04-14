package main

func climbStairs(n int) int {
	q, p, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

// f(x)=f(x−1.两数之和)+f(x−2),「滚动数组思想」
func main() {
	var c int = 5
	var nums int = climbStairs(c)
	println("Poblitily:", nums)
}
