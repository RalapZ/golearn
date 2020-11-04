package main

import (
	"fmt"
	"time"
)

func Worker(i int) chan int {
	a := make(chan int)
	go func() {
		//time.Sleep(time.Millisecond)
		for {
			fmt.Println(i, <-a)
		}
	}()
	return a
}

func main() {
	var a, b = Worker(0), Worker(1)
	go func() {
		for i := 0; i < 100; i++ {
			a <- i
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			b <- i
		}
	}()
	time.Sleep(time.Millisecond)
}
