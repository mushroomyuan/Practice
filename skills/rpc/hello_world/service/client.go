package service

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func NewClient() (HelloService, error) {
	// conn, err := rpc.Dial("tcp", "127.0.0.1:1234")
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{conn: client}, nil
}

type HelloServiceClient struct {
	conn *rpc.Client
}

func (c *HelloServiceClient) Hello(req *HelloRequest, resp *HelloResponse) error {
	return c.conn.Call("HelloService.Hello", req, resp)
}

func (c *HelloServiceClient) Close() error {
	return c.conn.Close()
}
