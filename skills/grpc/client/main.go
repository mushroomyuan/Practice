package main

import (
	"context"
	"fmt"

	"github.com/mushroomyuan/Practice/skills/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := service.NewHelloServiceClient(conn)
	response, err := client.SayHello(context.Background(), &service.HelloRequest{
		MyName: "yfz",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
