package main

import (
	"log"
	"micro-demo/grpc/demo"
	"micro-demo/grpc/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	demo.RegisterDemoServiceServer(grpcServer, new(service.DemoServer))
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
