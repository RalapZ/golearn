package main

import (
	"fmt"
	"time"
)

func Generate1() chan int {
	out := make(chan int)
	go func(out chan int) {
		n := 0
		for {
			//time.Sleep(800*time.Millisecond)
			out <- n
			n++
		}
	}(out)
	return out
}

func CreateW() chan int {
	in := make(chan int)

	go func() {
		for n := range in {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("received %d\n",
				n)
		}
	}()
	return in
}

func main() {
	c1, c2 := Generate1(), Generate1()
	var list []int
	//w:=CreateW()
	var w = CreateW()
	n := 0
	//hasvalue:=false
	for {
		var Cwork chan int
		var v int
		if len(list) > 0 {
			//n=list[0]
			v = list[0]
			Cwork = w
		}
		select {
		case n = <-c1:
			list = append(list, n)
			//hasvalue=true
		case n = <-c2:
			list = append(list, n)
			//hasvalue=true
		case Cwork <- v:
			list = list[1:]
			//hasvalue=false
		}
	}

}
