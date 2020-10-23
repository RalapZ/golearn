package main

import (
	"fmt"
	"runtime/debug"
)

func test2() {
	fmt.Println("=====")
	fmt.Println(string(debug.Stack()))
	debug.PrintStack()
}
func test1() {
	test2()
}

func main() {
	test1()
}
