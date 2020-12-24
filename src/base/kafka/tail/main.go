package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

var (
	filename = "./test.log"
	config   = tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
)

func main() {
	T, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println(err)
	}
	var k *tail.Line
	//var ok bool

	for true {
		//k,ok=<-T.Lines
		//if !ok{
		//	time.Sleep(100*time.Millisecond)
		//}
		//fmt.Println(k.Text)
		time.After(30 * time.Second)
		select {
		case k = <-T.Lines:
			fmt.Println(k.Text)
		default:
		}
	}
}
