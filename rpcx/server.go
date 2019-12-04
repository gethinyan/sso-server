package main

import (
	"flag"

	"github.com/gethinyan/sso-server/rpcx/services"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "172.16.21.71:8972", "server address")
)

// 注册所有服务
func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Auth", new(services.Auth), "")
	s.Serve("tcp", *addr)
}
