package main

import "fmt"

func Bubble1(list []int) {
	lenth := len(list)
	times := 0
	tag := false
	for i := 0; i < lenth-1; i++ {
		if tag == true {
			for j := 0; j < lenth-i-1; j++ {
				times++
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
					tag = true
				}
			}
		}
	}
	fmt.Println(list)
	fmt.Println(times)
}

func main() {
	list := []int{49, 4, 7, 1, 5, 7}
	Bubble1(list)
	lis := []int{1}
	Bubble1(lis)
	lis2 := []int{1, 2, 2, 2, 3, 5, 6, 8, 9}
	Bubble1(lis2)
}
