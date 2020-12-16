package main

import (
	"fmt"
	"net/rpc"
)

//type Client struct{}

type Params struct {
	Width, Heigth int
}

func main() {
	rep := 0
	C, _ := rpc.DialHTTP("tcp", ":9091")
	err := C.Call("Client.Cry", Params{1, 2}, &rep)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rep)
}
