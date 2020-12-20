package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	pb "microTest/helloworld/message"
)

func main() {
	S := micro.NewService(micro.Name("testtemp"))
	S.Init()
	hello := pb.NewGreeterService("testtemp", S.Client())
	resq, _ := hello.Hello(context.TODO(), &pb.Request{Name: "wode "})
	fmt.Println(resq)
}
