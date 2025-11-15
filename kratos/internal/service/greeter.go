package service

import (
	"context"

	v1 "kratos_bench/api/helloworld/v1"
	"kratos_bench/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

// Ping implements helloworld.GreeterServer.
func (s *GreeterService) Ping(ctx context.Context, in *v1.PingRequest) (*v1.PingReply, error) {
	return &v1.PingReply{Message: "pong"}, nil
}

// SayHelloName implements helloworld.GreeterServer.
func (s *GreeterService) SayHelloName(ctx context.Context, in *v1.HelloNameRequest) (*v1.HelloNameReply, error) {
	return &v1.HelloNameReply{Message: "my name is " + in.Name}, nil
}
