package rpctest

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Server struct {
	Conn net.Conn
}

func NewSession(conn net.Conn) *Server {
	return &Server{conn}
}

func (S *Server) Write(data []byte) error {
	buf := make([]byte, 4+len(data))
	fmt.Println(buf)
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	//S.Conn.Write()
	return nil
}
