package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
	proto "microTest/myfirstmicrointerface/proto"
)

func main() {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"10.10.8.150:2379",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("greeter"))

	service.Init()

	greeter := proto.NewRespInfoService("greeter", service.Client())

	rsp, err := greeter.RespInfo(context.TODO(), &proto.RalapRequest{Name: "莫雷"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
}
