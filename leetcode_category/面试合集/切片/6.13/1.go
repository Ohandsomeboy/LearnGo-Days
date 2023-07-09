package main

import "fmt"

func addNumber1(intSlince []int) []int {
	for _, v := range intSlince {
		v++
	}
	return intSlince
}

func addNumber2(intSlince []int) []int {
	for k := range intSlince {
		intSlince[k]++
	}
	return intSlince
}

func main() {
	//intSlince := []int{1, 1, 1}
	//res1 := addNumber1(intSlince)
	//fmt.Println(res1)
	//res2 := addNumber2(intSlince)
	//fmt.Println(res2)

	//slice := make([]int, 5, 10)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//
	//arr := []int{1, 2, 3, 4, 5}
	//
	//println(arr, slice)
	//slice = arr[1:] // 浅拷贝（共享内存）
	////slice := arr[1:] // 浅拷贝（共享内存）
	//
	//println(arr, slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))

	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:]
	// 2.3.4.5
	slice[0] = 1
	// 1.3.4.5
	// 1.1.3.4.5
	fmt.Println(arr)
	fmt.Println(slice)
}
