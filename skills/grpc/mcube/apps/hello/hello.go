package hello

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/v2/ioc"
	grpcserver "github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/mushroomyuan/Practice/skills/grpc/service"
	"google.golang.org/grpc"
	"io"
)

func init() {
	ioc.Controller().Registry(&HelloServiceServer{})
}

type HelloServiceServer struct {
	service.UnimplementedHelloServiceServer
	ioc.ObjectImpl
}

func (s *HelloServiceServer) Name() string {
	return "hello"
}

func (s *HelloServiceServer) Init() error {
	// 需要拿到全局的grpc的server对象

	service.RegisterHelloServiceServer(grpcserver.Get().Server(), &HelloServiceServer{})
	return nil
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
