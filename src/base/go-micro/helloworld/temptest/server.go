package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	pb "microTest/helloworld/message"
)

type Test struct{}

func (t *Test) Hello(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	resp.Greeting = req.Name + "my friends"
	return nil
}

func main() {
	S := micro.NewService(micro.Name("testtemp"))
	S.Init()
	pb.RegisterGreeterHandler(S.Server(), new(Test))
	S.Run()

}
