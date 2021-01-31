package main

import "fmt"

func InitHeap(s []int) {
	slen := len(s)
	for node := slen / 2; node >= 0; node-- {
		temp := node
		for temp < slen {
			var k int
			if 2*temp+1 >= slen {
				break
			} else {
				if s[temp] < s[2*temp+1] {
					s[temp], s[2*temp+1] = s[2*temp+1], s[temp]
					k = 2*temp + 1
				}
			}
			if 2*temp+2 >= slen {
				break
			} else {
				if s[temp] < s[2*temp+2] {
					s[temp], s[2*temp+2] = s[2*temp+2], s[temp]
					k = 2*temp + 2
				}
			}
			temp = k
		}
	}
	fmt.Println(s)
}
func main() {
	//s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s := []int{0}
	InitHeap(s)
	//fmt.Println(3/2)
	//fmt.Println(s)

}
