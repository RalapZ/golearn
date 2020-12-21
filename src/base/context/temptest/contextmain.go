package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func context2(ctx context.Context) {
Loop:
	for {
		fmt.Println("hello test")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
}
func contextmain(ctx context.Context) {
	defer wg2.Done()
	go context2(ctx)
Loop:
	for {
		fmt.Println("hello word")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
}

func main() {
	//var ctx context.Context
	ctx, cancel := context.WithCancel(context.Background())
	wg2.Add(1)
	go contextmain(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg2.Wait()

}
