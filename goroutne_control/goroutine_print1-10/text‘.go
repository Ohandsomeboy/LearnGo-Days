package main

import (
	"fmt"
	"sync"
)

func main() {

	//使用channel按顺序输出1-10
	wg := sync.WaitGroup{}
	wg.Add(10)
	chs := make(chan int)
	//通过一个全局变量控制进channel的顺序
	tag := 1
	for i := 1; i <= 10; i++ {
		go func(value int) {
			//死循环，保证按顺序进chan
			for {
				if tag == value {
					chs <- value
					break
				}
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Print(<-chs)
		wg.Done()
		tag++
	}
	wg.Wait()
}
