package main

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/mushroomyuan/Practice/skills/grpc/service"
	"google.golang.org/grpc"
)

type HelloServiceServer struct {
	service.UnimplementedHelloServiceServer
}

func (HelloServiceServer) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		Message: "hello, " + req.MyName,
	}, nil
}

func (HelloServiceServer) Chat(stream grpc.BidiStreamingServer[service.ChatRequest, service.ChatResponse]) error {
	for {
		// 接收一个请求
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println("收到信息为：", req.Message)
		stream.Send(&service.ChatResponse{
			Id:        req.Id,
			IsSuccess: true,
			Message:   "do somthing to" + req.Message,
		})
	}
}
func main() {
	server := grpc.NewServer()
	service.RegisterHelloServiceServer(server, &HelloServiceServer{})
	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(listenner); err != nil {
		panic(err)
	}
}
