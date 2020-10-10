package main

import (
	"fmt"
	"net/http"
)

func HttpOpen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}

func HttpServeStart(Port string) {
	http.HandleFunc("/list", HttpOpen)
	http.ListenAndServe(Port, nil)
}
func main() {
	HttpServeStart(":9000")
	//net.Addr()
}
