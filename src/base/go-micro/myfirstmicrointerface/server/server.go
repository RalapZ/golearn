package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	pb "microTest/myfirstmicrointerface/proto"
)

type Person struct{}

func (P *Person) RespInfo(ctx context.Context, in *pb.RalapRequest, out *pb.RalapResponse) error {
	out.Name = in.Name
	out.Age = 20
	out.Hobby = []string{"devops", "coding"}
	return nil
}

func main() {
	s := micro.NewService(micro.Name("test"))
	s.Init()
	pb.RegisterRespInfoHandler(s.Server(), new(Person))
	fmt.Println()
	s.Run()
}
