package main

import (
	"flag"

	"github.com/gethinyan/sso-server/rpcx/services"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

// 注册所有服务
func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(services.Arith), "")
	s.Serve("tcp", *addr)
}
