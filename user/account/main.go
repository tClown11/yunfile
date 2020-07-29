package main

import (
	"log"
	"user-server/handler"
	"user-server/proto"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("micro.service.user-server"),
	)

	//初始化相关操作
	service.Init()

	//注册服务实现的handler
	if err := proto.RegisterAccountServiceHandler(service.Server(), new(handler.AccountService)); err != nil {
		log.Print(err.Error())
		return
	}

	// 运行server
	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}
