package service

import (
	"fmt"
	"net/rpc"
)

func NewClient() (HelloService, error) {
	conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &HelloServiceClient{conn: conn}, nil
}

type HelloServiceClient struct {
	conn *rpc.Client
}

func (c *HelloServiceClient) Hello(req *HelloRequest, resp *HelloResponse) error {
	return c.conn.Call("HelloService.Hello", req, resp)
}
