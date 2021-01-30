package main

import "fmt"

func InitHeap(s []int) {
	len := len(s)
	for k := len/2 - 1; k >= 0; k-- {
		if s[k] < s[2*k+2] {
			s[2*k+2], s[k] = s[k], s[2*k+2]

		}
		if s[k] < s[2*k+1] {
			s[2*k+1], s[k] = s[k], s[2*k+1]
		}
	}
	fmt.Println(s)

}
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	InitHeap(s)
	//fmt.Println(3/2)
	//fmt.Println(s)

}
