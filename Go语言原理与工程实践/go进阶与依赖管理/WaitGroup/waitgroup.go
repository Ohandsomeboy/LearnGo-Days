package WaitGroup

import (
	"fmt"
	"sync"
	"time"
)

// 这个WaitGroup暴露了三个方法
// Add(delta int)、Done()、Wait()
// 计数器+delta, 计数器-1， 阻塞直到计数器为0

// 计数器：
// 开始协程+1; 执行结束-1; 主协程阻塞知道计数器为0;

func hello(i int) {
	println("hello goroutine:" + fmt.Sprint(i))
}

func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}

func ManyGoWait() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
