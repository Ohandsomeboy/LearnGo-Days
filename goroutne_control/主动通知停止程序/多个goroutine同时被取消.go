package 主动通知停止程序

import (
	"context"
	"fmt"
	"time"
)

func text4() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "node1")
	go worker(ctx, "node2")
	go worker(ctx, "node3")

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "got the stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}