package main

import (
	"fmt"
)

// 首先，通过使用make函数创建了三个channel，
// 分别是odd、even和quit。
//
// 然后，在第一个匿名的goroutine中，
// 使用循环将1到100之间的数字发送到odd和even的channel中。
// 如果数字是偶数，则发送到even，否则发送到odd。最后，
// 关闭odd和even的channel。
//
// 在第二个匿名的goroutine中，
// 使用select语句监听odd和even的channel。通过读取channel的值和是否成功接收到值的状态，
// 判断是从odd还是even的channel中接收到了值。然后，根据接收到的值打印相应的奇数或偶数。
// 如果接收操作失败（即channel已关闭），则通过向quit的channel发送true来通知程序退出。
//
// 最后，在主goroutine中，通过从quit的channel接收值来阻塞程序，直到接收到值为止。
// 这样可以确保在所有奇数和偶数都被打印完后程序才退出。
//
// 这段程序利用了goroutine和channel的并发特性，通过goroutine之间的通信来实现并发处理。
// 使用select语句可以实现非阻塞地监听多个channel，从而能够按照接收到的值的顺序打印奇数和偶数。

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
			if i%2 == 0 { // 发送有序，这样就保证接收到的也有序
				even <- i
			} else {
				odd <- i
			}
		}
		close(odd)
		close(even)
	}()

	go func() {
		for {
			select {
			case v, ok := <-odd:
				if ok {
					fmt.Println("Odd:", v)
				} else {
					quit <- true
					return
				}
			case v, ok := <-even:
				if ok {
					fmt.Println("Even:", v)
				} else {
					quit <- true
					return
				}
			}
		}
	}()
	<-quit
}

// 在第一个goroutine中，使用循环将1到100之间的数字发送到odd和even的channel中。
// 由于goroutine的执行是并发的，无法保证它们的执行顺序。
// 但由于发送操作是阻塞的，每次发送操作都会等待接收方接收后才会继续发送下一个数字。
// 因此，即使两个goroutine并发执行，发送的顺序仍然是有序的。
//
// 在第二个goroutine中，通过使用select语句监听odd和even的channel。
// select语句会选择其中一个可用的channel执行相应的操作。
// 由于两个channel是有序发送的，因此无论哪个channel的值先准备好，select语句都会先选择它，从而保证了按照顺序打印。
//
// 总结起来，通过使用有序发送和select语句的选择机制，这个程序能够保证按照1到100的顺序打印奇数和偶数。
