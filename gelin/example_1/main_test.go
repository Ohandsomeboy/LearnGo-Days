package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestExample(t *testing.T) {
	seed := int64(1234567890)
	rand.Seed(seed)

	nums := make([]int64, num)
	for i := 0; i < num; i++ {
		nums[i] = rand.Int63()
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	fmt.Println(max)
}
