package main

import (
	"github.com/micro/go-micro/web"
	"net/http"
)

func main() {

	server := web.NewService(web.Address(":8081")) // 路由
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go micro"))
	})
	_ = server.Run()
}
