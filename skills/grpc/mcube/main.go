package main

import (
	"github.com/infraboard/mcube/v2/ioc"
	grpcserver "github.com/infraboard/mcube/v2/ioc/config/grpc"
	_ "github.com/mushroomyuan/Practice/skills/grpc/mcube/apps/hello"
	"net"
)

func main() {
	// ioc初始化
	ioc.DevelopmentSetup()
	server := grpcserver.Get().Server()

	//service.RegisterHelloServiceServer(server, &hello.HelloServiceServer{})
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
