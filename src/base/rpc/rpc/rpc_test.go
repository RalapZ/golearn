package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWriter(t *testing.T) {
	// 定义地址
	addr := "127.0.0.1:8081"
	data1 := "hello"
	// 等待组定义
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 写数据的协程
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
		}
		s := Session{Conn: conn}
		err = s.Write([]byte(data1))
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 读数据的协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := Session{Conn: conn}
		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		// 最后一层校验
		if string(data) != data1 {
			t.Fatal(err)
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}
