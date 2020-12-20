package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	pb "microTest/myfirstmicrointerface/proto"
)

func main() {
	s := micro.NewService(micro.Name("test"))
	s.Init()
	clien := pb.NewRespInfoService("test", s.Client())
	resq, err := clien.RespInfo(context.TODO(), &pb.RalapRequest{Name: "TEST"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resq)
}
