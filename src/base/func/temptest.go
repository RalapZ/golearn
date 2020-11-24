package main

import "fmt"

type sum func(a, b int) (s int)

func myzone(a, b int) (s int) {
	return a + b
}

func test(Sum sum, a, b int) {
	fmt.Println(Sum(a, b))
}

func main() {
	//a=myzone
	test(myzone, 1, 2)
}
