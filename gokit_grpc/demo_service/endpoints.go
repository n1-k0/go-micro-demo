package demo_service

import (
	"context"
	"fmt"
	"micro-demo/gokit_grpc/demoGokit"

	"github.com/go-kit/kit/endpoint"
)

type DemoServiceRequest struct {
	a int
	b int
}

type DemoServiceResponse struct {
	c int
}

type DemoEndpoints struct {
	DemoEndpoint endpoint.Endpoint
}

func (s DemoEndpoints) Add(ctx context.Context, a int32, b int32) (int, error) {
	resp, err := s.DemoEndpoint(ctx, &demoGokit.InParams{
		A: a,
		B: b,
	})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Print("resp: ", resp)
	response := resp.(*demoGokit.Res)
	return int(response.C), err
}

func MakeDemoEndpoint(svc Service) endpoint.Endpoint {
	fmt.Print("makeendpo!!\n")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(DemoServiceRequest)
		a := req.a
		b := req.b
		res, _ := svc.Add(ctx, a, b)
		fmt.Println("makeend: ", res)
		return DemoServiceResponse{c: res}, nil
	}
}
