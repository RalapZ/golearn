package main

import (
	"fmt"
	"sync"
	"time"
)

var k bool = false
var wg sync.WaitGroup

func f() {
	defer wg.Done()
Loop:
	for {
		fmt.Println("hello word")
		time.Sleep(500 * time.Millisecond)
		if k {
			break Loop
		}
	}
}

func main() {

	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)
	k = true
	wg.Wait()

}
