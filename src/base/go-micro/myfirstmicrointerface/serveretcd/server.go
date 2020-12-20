package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	proto "microTest/myfirstmicrointerface/proto"
)

//#这里对应这里的代码  service Greeter {  #这个对象的
//#    rpc Hello(HelloRequest) returns (HelloResponse) {}  #这个拥有那些方法
//#}
type Greeter struct{}

func (g *Greeter) RespInfo(ctx context.Context, req *proto.RalapRequest, rsp *proto.RalapResponse) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

func main() {
	//#定义etcd的注册
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"10.10.8.150:2379",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("greeter")) //#这两部是关键，网上说的其他东西都可以不看，但两个必须要有

	service.Init()
	proto.RegisterRespInfoHandler(service.Server(), new(Greeter)) //protoc有多少个服务器，你就在这里写多少个

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
