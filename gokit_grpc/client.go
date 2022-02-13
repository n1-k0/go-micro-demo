package main

import (
	"context"
	"flag"
	"fmt"
	"micro-demo/gokit_grpc/demoGokit"
	service "micro-demo/gokit_grpc/demo_service"
	"time"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		fmt.Println("gRPC dial err:", err)
	}
	defer conn.Close()
	// fmt.Println("conn", conn)
	svr := NewStringClient(conn)
	result, err := svr.Add(ctx, 1, 2)
	if err != nil {
		fmt.Println("Check error", err.Error())
	}

	fmt.Println("result=", result)
}

func NewStringClient(conn *grpc.ClientConn) service.DemoEndpoints {

	var ep = grpctransport.NewClient(conn,
		"demoGokit.DemoService",
		"Add",
		DecodeDemoRequest,
		EncodeDemoResponse,
		demoGokit.Res{},
	).Endpoint()

	userEp := service.DemoEndpoints{
		DemoEndpoint: ep,
	}
	return userEp
}

func DecodeDemoRequest(ctx context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func EncodeDemoResponse(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}
