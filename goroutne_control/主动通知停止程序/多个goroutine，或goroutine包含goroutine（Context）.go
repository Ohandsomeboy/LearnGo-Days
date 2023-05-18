package 主动通知停止程序

import (
	"context"
	"fmt"
	"time"
)

func text3() {
	// context.Background() 表示为最上面的节点
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done(): // ctx.Done()回传一个channel的值回来
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}
