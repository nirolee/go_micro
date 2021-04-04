package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"new-micro/proto/hello"
)

type CapServer struct{}

func (c *CapServer) SayHello(ctx context.Context, req *hello.SayRequest, res *hello.SayResponse) error {
	res.Answer = "我们的口号是:" + req.Message + ""
	return nil
}

func main() {
	//创建服务
	service := micro.NewService(
		micro.Name("cap.hello.server"),
	)
	//初始化方法
	service.Init()

	hello.RegisterCapHandler(service.Server(), new(CapServer))
	//运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
