package service

import (
	"context"

	v1 "kratos_bench/api/helloworld/v1"
	"kratos_bench/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
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
	return &v1.HelloReply{Message: "my name is " + g.Hello}, nil
}

// Ping implements helloworld.GreeterServer.
func (s *GreeterService) Ping(ctx context.Context, in *emptypb.Empty) (*v1.PongReply, error) {
	return &v1.PongReply{Message: "pong"}, nil
}
