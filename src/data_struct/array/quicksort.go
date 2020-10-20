package main

import "fmt"

func quicksort(arr []int) []int {

	Len := len(arr)
	if Len == 0 {
		return arr
	}
	var temp int
	temp = arr[0]
	low := 0

	for high := Len - 1; high > low; high-- {
		if low == high {
			arr[low] = temp
			break
		}
		if temp > arr[high] {
			arr[low] = arr[high]
			low++
			for ; low < high; low++ {
				if arr[low] > temp {
					arr[high] = arr[low]
					break
				}
			}
		}
		arr[low] = temp
	}
	//fmt.Println(arr)
	quicksort(arr[:low])
	quicksort(arr[low+1:])
	return arr
}

func main() {
	a := []int{123, 4, 6578, 9, 23, 4, 6, 10, 8, 7, 5}
	b := quicksort(a)
	fmt.Println(b)
}
