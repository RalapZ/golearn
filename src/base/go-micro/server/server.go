package main

import (
	"context"
	pb "github.com/Ralap.Z/learn/src/base/go-micro/proto"
	"github.com/micro/go-micro"
)

type server struct{}

func (s *server) Zone(ctx context.Context, req *pb.ZoneRequest, resp *pb.ZoneResponse) error {

	resp.Name = req.Name
	resp.Hobby = []string{"coding", "devops"}
	return nil

}

func main() {
	service := micro.NewService(micro.Name("gomicro.zone"))
	service.Init()
	pb.RegisterZoneInterfaceHandler(service.Server(), &server{})
	service.Run()

}
