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

func InitMinHeap(s []int) {
	slen := len(s)
	for node := slen/2 - 1; node >= 0; node-- {
		temp := node
		for temp < slen {
			k := temp
			if 2*temp+1 > slen {
				break
			} else {
				if s[temp] > s[2*temp+1] {
					s[temp], s[2*temp+1] = s[2*temp+1], s[temp]
					k = 2*temp + 1
				}
			}
			if 2*temp+2 >= slen {
				break
			} else {
				if s[temp] > s[2*temp+2] {
					s[temp], s[2*temp+2] = s[2*temp+2], s[temp]
					k = 2*temp + 2
				}
			}
			if s[temp] < s[2*temp+2] && s[temp] < s[2*temp+1] {
				break
			}
			temp = k
		}
	}
	fmt.Println(s)
}
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//s := []int{0}
	InitMinHeap(s)
	//fmt.Println(3/2)
	//fmt.Println(s)

}
