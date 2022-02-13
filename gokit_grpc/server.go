package main

import (
	"context"
	"flag"
	"fmt"
	"micro-demo/gokit_grpc/demoGokit"
	service "micro-demo/gokit_grpc/demo_service"
	"net"
	"os"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc service.Service
	svc = service.ServiceImpl{}

	// add logging middleware
	// svc = service.LoggingMiddleware(logger)(svc)

	endpoint := service.MakeDemoEndpoint(svc)

	//创建健康检查的Endpoint
	// healthEndpoint := service.MakeHealthCheckEndpoint(svc)

	//把算术运算Endpoint和健康检查Endpoint封装至StringEndpoints
	endpts := service.DemoEndpoints{
		DemoEndpoint: endpoint,
	}

	handler := service.NewDemoServer(ctx, endpts)

	ls, _ := net.Listen("tcp", "127.0.0.1:8080")
	gRPCServer := grpc.NewServer()
	fmt.Print(111)
	demoGokit.RegisterDemoServiceServer(gRPCServer, handler)
	fmt.Print(222)
	gRPCServer.Serve(ls)

}
