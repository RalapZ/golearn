package main

import (
	"context"
	"fmt"
	hello "github.com/Ralap.Z/learn/src/base/testtemp/grpc/pbconf"
	"google.golang.org/grpc"
)

func main() {
	grpcconn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	clien := hello.NewPersonInfoClient(grpcconn)
	resq, _ := clien.PersonInfo(context.Background(), &hello.Request{Name: "myzone"})
	fmt.Println(resq.Hobby)
}
