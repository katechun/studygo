package main

import (
	"github.com/katechun/studygo/rpc/tools"
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "github.com/katechun/studygo/rpc/tools.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	RegisterHelloService(new(tools.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
