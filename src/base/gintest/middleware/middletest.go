package main

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type k func()

func Middle(v func()) k {
	return func() {
		t := time.Now()
		v()
		fmt.Println(time.Since(t))
	}
}

func test() {
	fmt.Println("test")
}

func main() {
	test := Middle(test)
	test()
}
