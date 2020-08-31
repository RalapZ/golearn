package main

import "fmt"

func fab(n int, sum int) int {
	if n == 0 {
		return sum
	}
	return fab(n-1, n+sum)
}

func main() {
	fmt.Println(fab(5, 0))
}
