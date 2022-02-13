package demo_service

import (
	"context"
	"fmt"
)

type Service interface {
	Add(ctx context.Context, a, b int) (int, error)
}

type ServiceImpl struct{}

func (s ServiceImpl) Add(ctx context.Context, a, b int) (int, error) {
	fmt.Printf("%d add %d: %d\n", a, b, a+b)
	return a + b, nil
}
