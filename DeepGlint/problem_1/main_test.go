package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestExample_1(t *testing.T) {
	//---------验证------------
	seed := int64(1234567890)
	rand.Seed(seed)

	start := time.Now()
	nums := make([]int64, num)
	for i := 0; i < num; i++ {
		nums[i] = rand.Int63()
	}
	elapsed := time.Since(start)

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	//-------------------------
}
func TestExample_2(t *testing.T) {
	seed := int64(1234567890)
	rand.Seed(seed)

	start := time.Now()
	nums := make([]int64, num)
	for i := 0; i < num; i++ {
		nums[i] = rand.Int63()
	}
	elapsed := time.Since(start)

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
		for localMax := range ch {
			if localMax > max {
				max = localMax
			}
		}
	}()
	wg.Wait()
	close(ch)

	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

func TestExample_3(t *testing.T) {
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
