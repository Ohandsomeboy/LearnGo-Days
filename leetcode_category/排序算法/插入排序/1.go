package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var length = 10
	var tree []int

	for i := 0; i < length; i++ {
		tree = append(tree, int(rand.Intn(1000)))
	}
	fmt.Println(tree)
	fmt.Println("---------------------")
	for i := 1; i < length; i++ {
		for j := i; j > 0 && tree[j] < tree[j-1]; j-- {
			tree[j], tree[j-1] = tree[j-1], tree[j]
		}
		fmt.Println(tree)
	}
}
