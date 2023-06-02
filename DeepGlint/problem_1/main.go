package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const num = 100000000 // 一亿随机数

func main() {
	seed := int64(1234567890)
	rand.Seed(seed)

	nums := make([]int64, num)
	for i := 0; i < num; i++ {
		nums[i] = rand.Int63()
	}

	max := nums[0]
	var wg sync.WaitGroup
	ch := make(chan int64)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			localMax := nums[(i-1)*num/10]
			for _, num := range nums[(i-1)*num/10 : i*num/10] {
				if num > localMax {
					localMax = num
				}
			}
			ch <- localMax
		}(i)
	}
	go func() {
		for i := 0; i < 9; i++ {
			localMax := <-ch
			if localMax > max {
				max = localMax
			}
		}
		close(ch)
	}()
	wg.Wait()

	fmt.Println(max)
}
