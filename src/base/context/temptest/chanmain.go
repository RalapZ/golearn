package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan bool)
var wg1 sync.WaitGroup

func test() {
	defer wg1.Done()
Loop:
	for {
		fmt.Println("hello word")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ch:
			break Loop
		default:
			continue
		}
	}
}

func main() {

	wg1.Add(1)
	go test()
	time.Sleep(5 * time.Second)
	ch <- true
	wg1.Wait()

}
