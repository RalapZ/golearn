package main

import "fmt"

func main() {
	//s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	slen := 10
	n := slen / 2
	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}
