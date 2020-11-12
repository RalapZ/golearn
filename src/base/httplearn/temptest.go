package main

import (
	"fmt"
	"log"
	"net/http"
)

type db map[string]int

func (db1 db) ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	log.Println(r)
	log.Fatal()
	for k, v := range db1 {
		fmt.Fprint(resp, k, v)
	}
}

func main() {
	db1 := db{"shoes": 5,
		"shift": 6}
	http.ListenAndServe(":9000", db1)

}
