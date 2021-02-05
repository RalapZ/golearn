//##################################################################################################//
//                   			         ┌─┐       ┌─┐ + +                                          //
//                   			      ┌──┘ ┴───────┘ ┴──┐++                                         //
//                   			      │       ───       │++ + + +                                   //
//                   			      ███████───███████ │+                                          //
//                   			      │       ─┴─       │+                                          //
//                   			      └───┐         ┌───┘                                           //
//                   			          │         │   + +                                         //
//                   			          │         └──────────────┐                                //
//                   			          │                        ├─┐                              //
//                   			          │                        ┌─┘                              //
//                   			          └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +                       //
//                   			            │ ─┤ ─┤       │ ─┤ ─┤                                   //
//                   			            └──┴──┘       └──┴──┘  + + + +                          //
//                   			      神兽出没               永无BUG                                 //
//   Author: Ralap                                                                                  //
//   Date  : 2021/01/31                                                                             //
//##################################################################################################//
package main

import "fmt"

func InitMaxHeap(s []int) {
	slen := len(s)
	for node := slen/2 - 1; node >= 0; node-- {
		temp := node
		k := node
		for temp < slen {
			//k:=temp
			if 2*temp+1 >= slen && 2*temp+2 >= slen {
				break
			} else if 2*temp+2 < slen {
				if s[temp] > s[2*temp+2] && s[temp] > s[2*temp+1] {
					break
				}
				if s[temp] < s[2*temp+2] {
					s[temp], s[2*temp+2] = s[2*temp+2], s[temp]
					k = 2*temp + 2
				}
				if s[temp] < s[2*temp+1] {
					s[temp], s[2*temp+1] = s[2*temp+1], s[temp]
					k = 2*temp + 1
				}
				temp = k

			} else {
				if s[temp] > s[2*temp+1] {
					break
				}
				if s[temp] < s[2*temp+1] {
					s[temp], s[2*temp+1] = s[2*temp+1], s[temp]
					k = 2*temp + 1
				}
			}
		}
	}
	fmt.Println(s)
}
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 234}
	//s := []int{8,7,6,5,4,3,0}
	InitMaxHeap(s)
	//fmt.Println(3/2)
	//fmt.Println(s)

}
