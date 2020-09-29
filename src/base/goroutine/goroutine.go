package main

import (
	"fmt"
	"time"
)

func GoPrint() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}

func main() {
	go GoPrint()
	time.Sleep(time.Millisecond)
}
