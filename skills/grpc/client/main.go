package main

import (
	"context"
	"fmt"
	"io"
	"sync"

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
	var wg sync.WaitGroup

	client := service.NewHelloServiceClient(conn)
	// response, err := client.SayHello(context.Background(), &service.HelloRequest{
	// 	MyName: "yfz",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(response)
	stream, err := client.Chat(context.Background())
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			fmt.Println(resp)
			wg.Done()
		}
	}()
	wg.Add(10)
	for i := range 10 {
		if err := stream.Send(&service.ChatRequest{
			Id:      int64(i),
			Message: fmt.Sprintf("hello %d", i),
		}); err != nil {
			fmt.Println(err)
			continue
		}
	}
	wg.Wait()
	fmt.Println("all finished!")

}
