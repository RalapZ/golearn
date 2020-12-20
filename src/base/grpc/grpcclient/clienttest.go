package main

import (
	"context"
	"fmt"
	pb "github.com/Ralap.Z/learn/src/base/grpc/pbconf"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithInsecure())
	defer conn.Close()
	resq := pb.NewResponseClient(conn)
	resp, e := resq.Myfirst(context.Background(), &pb.MyzoneReq{Myzonername: "test"})
	req, err := resp, e
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.Myzonermessage)
}
