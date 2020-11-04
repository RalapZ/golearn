package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) *
				time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func Worker1(id int, c chan int) {
	for a := range c {
		fmt.Println(id, ":", a)
	}
}

func CreateWorker1(id int) chan int {
	c := make(chan int)
	go Worker1(id, c)
	return c
}
func main() {
	a, b := Generate(), Generate()
	tm := time.After(time.Millisecond * 10000)
	var values []int
	worker := CreateWorker1(0)
	for {
		var activeWorker chan int
		var activeValue int
		//time.Sleep()
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-a:
			//fmt.Println("len :",len(a))
			values = append(values, n)
		case n := <-b:
			values = append(values, n)
			//fmt.Println("len:",len(values))
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tm:
			fmt.Println("time end")
			return
			//default:
			//	fmt.Println("no values")
		}
	}
}
