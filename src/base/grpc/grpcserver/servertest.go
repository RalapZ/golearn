package main

import (
	"context"
	"fmt"
	pb "github.com/Ralap.Z/learn/src/base/grpc/pbconf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Myzoneserver struct{}

func (S *Myzoneserver) Myfirst(ctx context.Context, in *pb.MyzoneReq) (*pb.MyzoneResp, error) {
	return &pb.MyzoneResp{Myzonermessage: "myfirstzone" + in.Myzonername}, nil
}

func main() {
	//
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterResponseServer(s, &Myzoneserver{})
	reflection.Register(s)
	s.Serve(lis)
}
