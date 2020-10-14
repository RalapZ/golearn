package main

import "fmt"

//type T struct{
//	v int
//}

func (k *int) add() {
	//k=&T{}
	*k = 1
	//return t
}

func main() {
	var t *int
	*t = 1
	fmt.Println(*t)
}
