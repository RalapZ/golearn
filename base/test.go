package main

import "fmt"

func main() {
	a := struct{}{}
	b := struct{}{}
	//if a == b {
	fmt.Printf("%p,%p", &a, &b)
	//}
}
