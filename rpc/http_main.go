package main

import (
	"github.com/katechun/studygo/rpc/tools"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}
	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
}

func main() {
	rpc.RegisterName("HelloService", new(tools.HelloService))

	http.Handle("/jsonrpc", &myHandler{})

	http.ListenAndServe(":1234", nil)
}
