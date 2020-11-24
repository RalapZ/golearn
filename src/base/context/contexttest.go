package main

import (
	"context"
	"log"
	"os"
	"time"
)

var (
	logg *log.Logger
)

func work(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			ch <- true
		default:
			log.Println("default")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	//ch := make(chan bool)
	//logg = log.New(os.Stdout, "", log.Ltime)
	//
	//ctx, cancel := context.WithCancel(context.Background())
	////context.
	//go work(ctx, ch)
	//time.Sleep(10 * time.Second)
	////取消函数：当cancel被调用时,WithCancel遍历Done以执行关闭；
	//cancel()
	//// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	//<-ch
	//logg.Println(`无脑发呆中!`)

	logg = log.New(os.Stdout, "", log.Ltime)
	ch := make(chan bool)
	cb := context.Background()
	ctx, cal := context.WithCancel(cb)
	go work(ctx, ch)
	time.Sleep(3 * time.Second)
	cal()
	<-ch
	log.Println("Done")
}
