package main

import (
	"fmt"
	"time"
)

type Stringslice []string

func (Stringslice) Len() int { return 0 }
func (Stringslice) Less(i, j int) bool {
	return false
}
func (Stringslice) Swap(i, j int) {}

func main() {
	for i := 0; i < 10; i++ {
		go func(k int) {
			fmt.Println(k)
		}(i)
	}
	time.Sleep(time.Millisecond)

}
