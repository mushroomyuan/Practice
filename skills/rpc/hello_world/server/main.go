package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 实现业务功能
// req := &HelloRequest{}
// resp := &HelloResponse{}
// err := &HelloServiceImpl{}.Hello(req, resp)
// net/rpc
// 1. 写好的对象， 注册给RPC Server
// 2. 再把RPC Server 启动起来
type HelloServiceImpl struct {
}

type HelloRequest struct {
	MyName string `json:"my_name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

// HTTP Handler
func (h *HelloServiceImpl) Hello(request *HelloRequest, response *HelloResponse) error {
	response.Message = "hello:" + request.MyName
	return nil
}

func main() {
	// 1.把服务对象注册到rpc框架
	if err := rpc.RegisterName("HelloService", &HelloServiceImpl{}); err != nil {
		fmt.Println(err)
		return
	}
	// 2.启创建监听对象
	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listenner.Close()
	// 3.启动监听，循环接受消息
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 4.将接收的消息发送给rpc框架
		go rpc.ServeConn(conn)
	}

}
