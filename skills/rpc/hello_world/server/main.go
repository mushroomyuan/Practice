package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/mushroomyuan/Practice/skills/rpc/hello_world/service"
)

type HelloServer struct {
}

func (h *HelloServer) Hello(request *service.HelloRequest, response *service.HelloResponse) error {
	response.Message = "hello:" + request.MyName
	return nil
}

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func main() {
	// 1.把服务对象注册到rpc框架
	if err := rpc.RegisterName("HelloService", &HelloServer{}); err != nil {
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

		// go rpc.ServeConn(conn)
		// json格式的编解码器：
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //
	}

	// http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {

	// 	conn := NewRPCReadWriteCloserFromHTTP(w, r)
	// 	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	// })
	// http.ListenAndServe(":1234", nil)
}
