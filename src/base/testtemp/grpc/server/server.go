package main

import (
	"context"
	pb "github.com/Ralap.Z/learn/src/base/testtemp/grpc/pbconf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct{}

func (s *Server) PersonInfo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	resp := new(pb.Response)
	resp.Name = req.Name
	resp.Age = 20
	resp.Hobby = []string{"devops", "coding"}
	return resp, nil
}
func main() {
	lis, _ := net.Listen("tcp", "127.0.0.1:8080")
	server := grpc.NewServer()
	pb.RegisterPersonInfoServer(server, &Server{})
	reflection.Register(server)
	server.Serve(lis)

}
