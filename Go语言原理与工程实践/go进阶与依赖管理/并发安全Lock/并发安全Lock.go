package main

import (
	"sync"
	"time"
)

// 对变量执行2000次+1操作
// 5个协程并发执行,预期结果是10000

var (
	x    int
	lock sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("WithoutLock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("WithLock:", x)
}

func main() {
	Add()
}
