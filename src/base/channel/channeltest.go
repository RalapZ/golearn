package main

import "fmt"

func main() {
	test := make(chan int, 10)
	Check := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println("test",i)
			test <- i
		}

	}()
	fmt.Println(len(test), cap(test))
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-test)
		}
		Check <- true
	}()
	<-Check
}
