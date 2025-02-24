package main

import (
	"context"
	"net"

	"github.com/mushroomyuan/Practice/skills/grpc/service"
	"google.golang.org/grpc"
)

type UnimplementedHelloServiceServer struct {
	service.UnimplementedHelloServiceServer
}

func (UnimplementedHelloServiceServer) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		Message: "hello, " + req.MyName,
	}, nil
}
func main() {
	server := grpc.NewServer()
	service.RegisterHelloServiceServer(server, &UnimplementedHelloServiceServer{})
	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(listenner); err != nil {
		panic(err)
	}
}
