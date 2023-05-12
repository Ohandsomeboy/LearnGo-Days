package main

import (
	"math/rand"
	"sync"
)

const N = 100000
const M = 2

var s [N]int
var mutexes [N]sync.Mutex

func updateS() {
	for i := 0; i < 10000; i++ {
		// 生成随机的i和jl
		i := rand.Intn(N)
		j := rand.Intn(N)

		// 获取i, i+1, i+2的锁
		mutexes[i].Lock()
		mutexes[(i+1)%N].Lock()
		mutexes[(i+2)%N].Lock()

		// 读取i, i+1, i+2的值
		x := s[i]
		y := s[(i+1)%N]
		z := s[(i+2)%N]

		// 更新s[j]
		s[j] = x + y + z

		// 释放i, i+1, i+2的锁
		mutexes[i].Unlock()
		mutexes[(i+1)%N].Unlock()
		mutexes[(i+2)%N].Unlock()
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
