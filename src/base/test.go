package main

import (
	"fmt"
	"os"
)

func main() {
	a := struct{}{}
	b := struct{}{}
	//if a == b {
	fmt.Printf("%p,%p", &a, &b)
	k, err := fmt.Printf("%p,%p\n", &a, &b)
	fmt.Fprintln(os.Stdout, k, err)
	//fmt.Sprint
	//}
}
