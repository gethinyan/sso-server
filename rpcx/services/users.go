package services

import (
	"context"
)

// Args 请求参数结构
type Args struct {
	A int
	B int
}

// Reply 响应参数结构
type Reply struct {
	C int
}

// Arith ...
type Arith int

// Mul 注册函数
func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}
