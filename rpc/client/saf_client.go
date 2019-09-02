package main

import "log"
import "net/rpc"

const HelloServiceName = "github.com/katechun/studygo/rpc/tools.HelloService"

type HelloServiceClient struct {
	*rpc.Client
}

var HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloSercice(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}
