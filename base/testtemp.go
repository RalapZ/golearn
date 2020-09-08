package main

import "fmt"

func main() {
	a := "tesasdfasdfasdfdfasd"
	b := a[len("tes"):]
	fmt.Println(b)
}
