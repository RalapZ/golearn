package main

import (
	"fmt"
	"time"
)

func test(a chan int) {
	i := 0
	for {
		a <- i
		i++
		fmt.Println("test")
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	a := make(chan int)
	var b chan int
	go test(a)
	t := time.After(1000 * time.Millisecond)
L:
	for {
		//j:=0
		select {
		case <-a:
		case <-b:
		case <-t:
			//j=1
			break L
		}
		//if j!=0{
		//	break
		//}
	}
	fmt.Println("Done")
}
