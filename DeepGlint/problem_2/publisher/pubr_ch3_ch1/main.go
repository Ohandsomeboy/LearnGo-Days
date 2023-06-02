package main

import (
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
)

// 从命令行读取整数
func readIntFromCommandLine(reader *bufio.Reader, target *string) *string {
	_, err := fmt.Fscanf(reader, "%s\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
	return target
}

func main() {
	// 创建一个新的Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for {
		fmt.Println("Please input Message on to Terminal")
		var stdinStr string
		reader := bufio.NewReader(os.Stdin)
		readIntFromCommandLine(reader, &stdinStr)

		// 将消息发布到名为 channel_1 和 channel_3 的Redis通道中
		if err := client.Publish("channel_1", stdinStr).Err(); err != nil {
			panic(err)
		}
		err := client.Publish("channel_3", stdinStr).Err()
		if err != nil {
			panic(err)
		}
	}
}
