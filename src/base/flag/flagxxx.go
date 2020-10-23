package main

import (
	"flag"
	"fmt"
)

var (
	test int
)

func init() {
	flag.IntVar(&test, "flag", 1, "my first int")
}
func main() {
	flag.Parse()
	fmt.Println(test)
}
