package main

import (
	"fmt"
	"net/http"
)

func Catfile(writer http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(writer, "hello")
}

func main() {
	http.HandleFunc("/hello", Catfile)
	http.ListenAndServe(":8190", nil)
}
