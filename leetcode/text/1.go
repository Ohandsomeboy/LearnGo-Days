package main

import "fmt"

func main() {
	var a = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:3]
	fmt.Println(b)
	b = append(b, 11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println(len(b), cap(b))
	fmt.Println(len(a), cap(b))
	fmt.Println(a, b)
}
