package demo_service

import (
	"context"
	"fmt"
	"micro-demo/gokit_grpc/demoGokit"

	"github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	add grpc.Handler
}

func (s *grpcServer) Add(ctx context.Context, r *demoGokit.InParams) (*demoGokit.Res, error) {
	_, resp, err := s.add.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*demoGokit.Res), nil
}

func NewDemoServer(ctx context.Context, endpoint DemoEndpoints) demoGokit.DemoServiceServer {
	return &grpcServer{
		add: grpc.NewServer(
			endpoint.DemoEndpoint,
			DecodeDemoServiceRequest,
			EncodeDemoServiceResponse,
		),
	}
}

func DecodeDemoServiceRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*demoGokit.InParams)
	fmt.Print("decode: ", req)
	return DemoServiceRequest{a: int(req.A), b: int(req.B)}, nil
}

func EncodeDemoServiceResponse(ctx context.Context, r interface{}) (interface{}, error) {
	fmt.Println("encoe: ", r)
	resp := r.(DemoServiceResponse)
	fmt.Print("encode: ", resp)
	return &demoGokit.Res{
		C: int32(resp.c),
	}, nil
	// return DemoServiceResponse{c: int(resp.c)}, nil
}
