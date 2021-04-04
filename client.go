package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"new-micro/proto/hello"
)

func main() {
	service := micro.NewService(
		micro.Name("cap.hello.client"),
	)
	//初始化
	service.Init()
	capHello := hello.NewCapService("cap.hello.server", service.Client())
	res, err := capHello.SayHello(context.TODO(), &hello.SayRequest{
		Message: "如果不是真的废话，谁又愿意混吃等死呢",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Answer)
}
