package main

import "fmt"

func Bubble(list []int) {
	lenth := len(list)
	times := 0
	for i := 0; i < lenth-1; i++ {
		for j := 0; j < lenth-i-1; j++ {
			times++
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	fmt.Println(list)
	fmt.Println(times)
}

func main() {
	list := []int{49, 4, 7, 1, 5, 7}
	Bubble(list)
	lis := []int{1}
	Bubble(lis)
	lis2 := []int{1, 2, 2, 2, 3, 5, 6, 8, 9}
	Bubble(lis2)

}
