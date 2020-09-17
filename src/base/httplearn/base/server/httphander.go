package server

import (
	"fmt"
	"net/http"
)

type Hello struct {
	content string
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.content)
}

func main() {
	H := Hello{"hello my handle test"}
	http.Handle("/", &H)
	fmt.Printf("%T", H)
	http.ListenAndServe(":8090", nil)
}
