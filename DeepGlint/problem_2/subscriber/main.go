package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

func main() {
	// 创建一个新的Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 订阅名为“mychannel”的Redis通道
	pubsub := client.Subscribe("channel_1", "channel_2", "channel_3")
	defer pubsub.Close()

	fmt.Println("Subscribed to channel_1， channel_2， channel_3")

	for {
		// 从标准输入中读取数字按键
		fmt.Print("Enter a number key (1, 2 or 3): ")
		reader := bufio.NewReader(os.Stdin)
		key, _ := reader.ReadString('\n')
		key = strings.TrimSpace(key)

		// 检查数字按键是否有效
		i, err := strconv.Atoi(key)
		if err != nil || i < 1 || i > 3 {
			fmt.Println("Invalid number key. Please enter 1, 2 or 3.")
			continue
		}

		// 接收指定的频道消息
		msgi := i - 1
		ch := pubsub.Channel()
		for msg := range ch {
			if msgi == 0 && msg.Channel == "channel_1" {
				fmt.Printf("Received message from channel %s: %s\n", msg.Channel, msg.Payload)
				if msg.Payload == "done" {
					break
				}
			} else if msgi == 1 && msg.Channel == "channel_2" {
				fmt.Printf("Received message from channel %s: %s\n", msg.Channel, msg.Payload)
				if msg.Payload == "done" {
					break
				}
			} else if msgi == 2 && msg.Channel == "channel_3" {
				fmt.Printf("Received message from channel %s: %s\n", msg.Channel, msg.Payload)
				if msg.Payload == "done" {
					break
				}
			}
		}
	}
}
