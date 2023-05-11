package main

import (
	"math/rand"
	"sync"
)

const N = 100000
const M = 10000

var s [N]int
var mutex sync.RWMutex

func updateS() {
	for i := 0; i < M; i++ {
		// 生成随机的i和j
		i := rand.Intn(N)
		j := rand.Intn(N)

		// 获取写锁
		mutex.Lock()

		// 读取s(i), s(i+1), s(i+2)
		x := s[i]
		y := s[(i+1)%N]
		z := s[(i+2)%N]

		// 更新s(j)
		s[j] = x + y + z

		// 释放写锁
		mutex.Unlock()
	}
}

func main() {
	// 初始化数组s...

	// 创建并发工作者
	var wg sync.WaitGroup
	for i := 0; i < M; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateS()
		}()
	}

	// 等待所有工作者完成
	wg.Wait()

	// 打印结果...
}
