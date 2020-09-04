package main

import "fmt"

func deferfunc(){
	defer fmt.Println("test")
	panic("error print")
	defer fmt.Println("tes1")

}
func main() {
	deferfunc()
}
