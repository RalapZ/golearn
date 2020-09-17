package server

import (
	"fmt"
	"net/http"
)

func Sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my first web server")
}
func main() {
	http.HandleFunc("/", Sayhello)
	http.ListenAndServe(":8000", nil)
}
