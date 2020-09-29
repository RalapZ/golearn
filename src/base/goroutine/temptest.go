package main

import (
	"time"
)

func main() {
	//var a int
	for i := 0; i < 100; i++ {
		go func(ii int) {
			for {
				ii++

			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Print(i)
}
