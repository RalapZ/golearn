package main

import (
	"net/http"
	"net/rpc"
)

type Client struct{}

type Params struct {
	Width, Heigth int
}

func (C *Client) Cry(P Params, res *int) error {
	*res = (P.Heigth + P.Width) * 2
	return nil
}

func main() {
	ract := new(Client)
	rpc.Register(ract)
	rpc.HandleHTTP()
	http.ListenAndServe(":9091", nil)
}
