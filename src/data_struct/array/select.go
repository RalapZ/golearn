package main

import "fmt"

func SelectFun(list []int) {
	//var lenth int
	lenth := len(list)
	for i := 0; i < lenth; i++ {
		for j := i; j < lenth; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println(list)
}

func main() {
	lis := []int{1, 2, 3, 4, 5, 6, 2, 3}
	SelectFun(lis)
}
