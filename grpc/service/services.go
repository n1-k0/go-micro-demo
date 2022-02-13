package service

import (
	"context"
	"fmt"
	"micro-demo/grpc/demo"
)

type DemoServer struct{}

func (s *DemoServer) Demo(ctx context.Context, req *demo.String) (*demo.String, error) {
	fmt.Printf("get req.Value: %s", req.Value)
	req.Value = "2"
	return req, nil
}


