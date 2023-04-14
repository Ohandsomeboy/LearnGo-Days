package main

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if m[v] == 0 {
			m[v]++
		} else if m[v] == 1 {
			delete(m, v)
		}
	}
	ans := 0
	for i, v := range m {
		if v != 0 {
			return i
		}
	}
	return ans
}

// 利用任何数和0做异或运算，a异或0=a
// 任何数和其自身做异或运算，a异或a=0
// 异或运算满足交换律和结合律
//func singleNumber(nums []int) int {
//	single := 0
//	for _, num := range nums {
//		single ^= num
//	}
//	return single
//}

func main() {

}
