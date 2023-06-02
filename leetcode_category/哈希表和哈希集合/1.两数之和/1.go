package main

import "fmt"

// 时间复杂度：O(N^2) 空间复杂度：O(1.两数之和)
//func twoSum(nums []int, target int) []int {
//	for i, x := range nums {
//		for j := i + 1.两数之和; j < len(nums); j++ {
//			if x+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

// ok为True，false，p为先前存进hashmap里的nums位置
// 每次都存进去查询哈希表内有无满足target-x的值，有就输出，无就继续
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums { //i是迭代，x为值value，一开始为nums[0] == x,0 = i
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil

	//m := map[int]int{}
	//for i,x := range nums {
	//	for p,ok := m[target-x]; ok{
	//		return []int{p,i}
	//	}
	//	m[x]=i
	//}
	//return nil
}

func main() {
	var num_array = []int{2, 7, 11, 15, 31, 9}
	var c []int
	c = twoSum(num_array, 16)
	fmt.Println(c)
}
