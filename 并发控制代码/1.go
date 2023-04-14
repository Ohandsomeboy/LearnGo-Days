package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// 模拟工作
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d working...\n", id)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}

// 这个例子中，我们使用了sync.WaitGroup来等待三个goroutine完成。在worker函数中，我们使用defer语句来确保Done方法总是被调用。
// 在main函数中，我们使用Add方法来增加WaitGroup的计数器，
// 然后在每个goroutine完成时调用Done方法来减少计数器。最后，我们调用Wait方法来阻塞main函数，直到所有的goroutine都完成1。
