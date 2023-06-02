package main

import (
	"context"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"time"

	pb "github.com/HanFa/learn-go/grpc-example/route"
	"google.golang.org/grpc"
)

// 这里内嵌UnimplementedXXX的原因：因为 protoc 在解析.proto文件，
// 创建grpc.pb.go 文件的时候用unplementedXXX 完成向前兼容的实现
type routeGuideServer struct {
	features []*pb.Feature // 做一个demo的数据库，指像pb.Feature的slice
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	// 用一个最简单的loop来遍历db
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) { // proto.Equal()可以比较两个Message是否相同
			return feature, nil
		}
	}
	return nil, nil
}

func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

func (s *routeGuideServer) ListFeatures(rectangle *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.features {
		if inRange(feature.Location, rectangle) { // 如果在这个Rectangle里，就把它放回返回流
			if err := stream.Send(feature); err != nil { // 这里使用流的原因：如果数据很多，就可以一边发送一边处理。就不需要一块传完再发送
				return err
			}
		}
	}
	return nil
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

func calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000)
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// RecordRoute 需要获取我拿到了几个点，走了多少距离(千米)，经过的时间（秒）
func (s *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	startTime := time.Now()
	var pointCount, distance int32
	var prevPoint *pb.Point
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			// conclude a route summary
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:  pointCount,
				Distance:    distance,
				ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
			}) // 发送最后一个东西，并把这个steam关闭
		}
		if err != nil {
			return err
		}
		pointCount++
		if prevPoint != nil {
			distance += calcDistance(prevPoint, point)
		}
		prevPoint = point
	}
	return nil
}

// 对一个请求进行Recommend，遍历features，根据用户的需求返回（数据大用priority queue）
func (s *routeGuideServer) recommendOnce(request *pb.RecommendationRequest) (*pb.Feature, error) {
	var nearest, farthest *pb.Feature
	var nearestDistance, farthestDistance int32

	for _, feature := range s.features {
		distance := calcDistance(feature.Location, request.Point)
		if nearest == nil || distance < nearestDistance {
			nearestDistance = distance
			nearest = feature
		}
		if farthest == nil || distance > farthestDistance {
			farthestDistance = distance
			farthest = feature
		}
	}
	if request.Mode == pb.RecommendationMode_GetFarthest {
		return farthest, nil
	} else {
		return nearest, nil
	}
}

func (s *routeGuideServer) Recommend(stream pb.RouteGuide_RecommendServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF { // 如果客户发送完了
			return nil
		}
		if err != nil {
			return err
		}
		recommended, err := s.recommendOnce(request)
		if err != nil {
			return err
		}
		stream.Send(recommended)
	}
}

// 这里的 RouteGuideServer 是个接口类型，需要一个type来实现
func newServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{ // 初始化数据库
			{Name: "上海交通大学闵行校区 上海市闵行区东川路800号", Location: &pb.Point{
				Latitude:  310235000,
				Longitude: 121437403,
			}},
			{Name: "复旦大学 上海市杨浦区五角场邯郸路220号", Location: &pb.Point{
				Latitude:  312978870,
				Longitude: 121503457,
			}},
			{Name: "华东理工大学 上海市徐汇区梅陇路130号", Location: &pb.Point{
				Latitude:  311416130,
				Longitude: 121424904,
			}},
		},
	}
}

func main() {
	// 新建一个network listener
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}

	// 新建一个grpc的server
	// NewServer creates a gRPC server which has no service registered and hos not
	// started to accept requests yet.
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer()) // 把在这个 proto 定义的服务，注册给新建的 grpcServer
	log.Fatalln(grpcServer.Serve(lis))                   // 这里就开启服务，等待call其中的一个method就可以对他进行响应
}
