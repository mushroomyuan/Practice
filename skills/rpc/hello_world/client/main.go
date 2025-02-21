package main

import (
	"fmt"

	"github.com/mushroomyuan/Practice/skills/rpc/hello_world/service"
)

func main() {
	client, err := service.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &service.HelloRequest{
		MyName: "bob",
	}
	resp := &service.HelloResponse{}

	if err := client.Hello(req, resp); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
