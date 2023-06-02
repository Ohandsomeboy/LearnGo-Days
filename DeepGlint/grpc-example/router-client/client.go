package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/HanFa/learn-go/grpc-example/route"
	"google.golang.org/grpc"
)

// 验证里客户端可以从服务端那里获取信息并产生响应的请求
func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000, // 这里是将经度和维度乘以10的七次方，将它变成int32类型
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

// 给一个Rectangle，然后返回这个 Rectangle 里面的 Feature
func runSecond(client pb.RouteGuideClient) {
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{ // 地图上两个点，然后可以包含两所大学的经纬度
		Lo: &pb.Point{Latitude: 313374060, Longitude: 121258540},
		Hi: &pb.Point{Latitude: 311074130, Longitude: 121598790},
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		feature, err := serverStream.Recv()
		if err == io.EOF { // 如果这个流关闭，会得到一个EOF(正常的)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}

}

// (客户端：每隔一秒上传一个点)
func runThird(client pb.RouteGuideClient) {
	// fummy data
	points := []*pb.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}

	clientStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// 遍历这些点，然后通过这个clientStream上传给后端
	for _, point := range points {
		if err := clientStream.Send(point); err != nil { // 注意处理handle error
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}

	// (Tip：在这里不能像上面那样写if再err前面，summary会调不到)
	summary, err := clientStream.CloseAndRecv() // 最后返回一个路线图总结
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(summary)
}

// 从标准流里读一个integer，然后存在这个target里面
func readIntFromCommandLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

// 客户端不断的发送询问，我这个点在这，距离这个点最近的Feature在哪，然后返回这个Feature
func runForth(client pb.RouteGuideClient) {
	// 这个steam是个双向通道
	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// 这个goroutine 监听 server steam
	go func() {
		for {
			feature, err2 := stream.Recv()
			if err2 != nil {
				log.Fatalln(err2)
			}
			fmt.Println("Recommended:", feature)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	// 在这里实现一个命令行的shell，通过输入经纬度向服务器询问最近或最远的Feature
	for {
		request := pb.RecommendationRequest{Point: new(pb.Point)} // 在这里新建一个空的point，不新point.Latitude会报错
		var mode int32
		fmt.Print("Enter Recommendation Mode (0 for farthest, 1 for the nearest)")
		readIntFromCommandLine(reader, &mode)
		fmt.Print("Enter Latitude:")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude:")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)

		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(100 * time.Millisecond) // 防止上面这个打印出来的字符Overlap
	}
}

func main() {
	// 1.访问的地址;	 2.忽略这个证书验证;	 3.将这dial变成这个blocking，拨号成功的才继续往下执行；
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server.")
	}
	defer conn.Close()

	// RouteGuide:对应这个服务的 method
	client := pb.NewRouteGuideClient(conn)

	// 1.Unary
	//runFirst(client)

	// client-side streaming
	//runSecond(client)

	// server-side streaming
	//runThird(client)

	// bidrectional streaming
	runForth(client)
}
