package main

import (
	"fmt"
	"sort"
)

type Stringslice []string

func (Strs Stringslice) Len() int {
	return len(Strs)
}
func (Strs Stringslice) Less(i, j int) bool {
	return Strs[i] < Strs[j]
}
func (Strs Stringslice) Swap(i, j int) {
	Strs[i], Strs[j] = Strs[j], Strs[i]
}

func main() {
	Strs := Stringslice{"d", "a", "c", "b", "gg", "f", "rr", "ww"}
	sort.Sort(Strs)
	fmt.Println(Strs)

}
