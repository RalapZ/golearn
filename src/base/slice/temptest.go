package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%p,%v\n", &a, a)
	fmt.Printf("%p,%v\n", &a, a)
	b := a[1:]
	fmt.Printf("%p,%v\n", &b, b)
	for i := 0; i < 5; i++ {
		b = append(b, 1)
		fmt.Printf("%p,%v\n", &b, b)
		fmt.Println(cap(b), len(b))
	}
}
