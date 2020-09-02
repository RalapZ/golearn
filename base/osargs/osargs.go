package main

import (
	"fmt"
	"os"
)

func main() {
	test := os.Args
	fmt.Println(len(test))
	fmt.Println(test)
}
