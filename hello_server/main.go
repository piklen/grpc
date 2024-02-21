package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello_server/pb"
	"net"
)

// hello server

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}
func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	fmt.Printf("创建gRPC服务器")
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
