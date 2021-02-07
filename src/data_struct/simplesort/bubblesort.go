package main

import "fmt"

func BubbleSort(k []int) {
	len := len(k)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if k[j] > k[j+1] {
				k[j], k[j+1] = k[j+1], k[j]
			}
		}

	}

	for _, v := range k {
		fmt.Println(v)
	}
}

func main() {
	//a:=[]int{1,2,4,4,5,1,2,56,234}
	a := []int{}
	BubbleSort(a)
}
