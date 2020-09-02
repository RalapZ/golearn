package main

import "fmt"

func binarysearch(n int, array []int) interface{} {
	if n < array[0] || n > array[len(array)-1] {
		return "error"
	}
	var f, b int
	f = 0
	b = len(array) - 1
	half := 0
	//k:=0
	for {
		fmt.Println(f, b)
		half = len(array[f:b])
		if array[half] < n {
			f = half + 1
		} else if array[half] > n {
			b = half
		} else if array[half] == n {
			return half
		} else if f <= b {
			return "not exist"
		}
	}
}
func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	n := 100
	fmt.Println(binarysearch(n, array))
}
