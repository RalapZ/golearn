package rpctest

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestServer_WriteRead(t *testing.T) {
	addr := "127.0.0.1:18081"
	data := "Ralap"
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		Lis, err := net.Listen("tcp", addr)
		defer Lis.Close()
		if err != nil {
			fmt.Println(err)
		}
		conn, err := Lis.Accept()
		s := NewSession(conn)
		_ = s.Write([]byte(data))
	}()
	go func() {
		wg.Done()
		conn, err := net.Dial("tcp", addr)
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
		}
		s := NewSession(conn)
		//noinspection GoUnhandledErrorResult
		data, err := s.Read()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}
