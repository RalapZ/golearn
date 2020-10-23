package main

import (
	"flag"
	"fmt"
)

var (
	test1 *int
)

func init() {
	test1 = flag.Int("flag1", 0, "you know")

}

func main() {
	flag.Parse()
	fmt.Println(*test1)
}
