package main

import (
	"fmt"
	"net/rpc"
)

type HelloRequest struct {
	MyName string `json:"my_name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	req := &HelloRequest{
		MyName: "bob",
	}
	resp := &HelloResponse{}

	err = conn.Call("HelloService.Hello", req, resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
