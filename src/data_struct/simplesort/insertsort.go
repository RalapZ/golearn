package main

import "fmt"

func InsertSort(a []int) {
	len := len(a)
	for i := 0; i < len; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	for i := 0; i < len; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	a := []int{1, 2, 4, 4, 5, 1, 2, 56, 234}
	InsertSort(a)
}
