package main

import (
	"fmt"
)

func Recoverytest() {
	defer func() {
		a := recover()
		fmt.Println(a)
	}()
	panic("asdfasdfasdfasdfsd")
}

func main() {
	Recoverytest()
}
